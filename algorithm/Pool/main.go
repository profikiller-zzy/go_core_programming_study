package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type Task func()

type Pool struct {
	taskChan       chan Task      // 存储任务的队列
	maxWorkerCount int            // 工人最大值
	workerCount    int            // 正在运行的工人
	idleTimeout    time.Duration  // 如果一个工人在idleTimeout时间内没有接到任务，则它自己退出
	mu             sync.Mutex     // 保护结构体中的字段
	wg             sync.WaitGroup // 确保所有工人成功退出
	closed         bool           // 该协程池是否已经关闭
}

func NewPool(queueSize int, maxWorkerCount int, idleTimeout time.Duration) *Pool {
	if maxWorkerCount <= 0 {
		maxWorkerCount = 1
	}
	if queueSize <= 0 {
		queueSize = 20
	}
	return &Pool{
		taskChan:       make(chan Task, queueSize),
		maxWorkerCount: maxWorkerCount,
		idleTimeout:    idleTimeout,
	}
}

func (p *Pool) Submit(task Task) error {
	p.mu.Lock()
	if p.closed {
		p.mu.Unlock()
		return errors.New("pool is closed")
	}
	p.taskChan <- task
	if p.workerCount < p.maxWorkerCount && p.workerCount < len(p.taskChan) {
		p.StartWorker()
	}
	p.mu.Unlock()
	return nil
}

func (p *Pool) SetParallelism(n int) error {
	if n <= 0 {
		return errors.New("invalid parallelism")
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	p.maxWorkerCount = n
	for p.workerCount < p.maxWorkerCount && p.workerCount < len(p.taskChan) {
		p.StartWorker()
	}
	return nil
}

// StartWorker 开启一个工人协程，调用该方法时必须持有锁
func (p *Pool) StartWorker() {
	p.wg.Add(1)
	p.workerCount++
	go func() {
		defer func() {
			p.wg.Done()
			p.mu.Lock()
			p.workerCount--
			p.mu.Unlock()
		}()
		timer := time.NewTimer(p.idleTimeout)
		defer timer.Stop()
		for {
			select {
			case task, ok := <-p.taskChan:
				if !ok {
					return
				}
				task() // 运行任务
				// 重制该计时器
				if !timer.Stop() {
					// 这里一定要用select + default清空timer，因为该timer可能会由于竞争
					// 让timer.Stop()返回false同时timer.C中没有值，这种写法防止程序永久阻塞
					select {
					case <-timer.C:
					default:
					}
				}
				timer.Reset(p.idleTimeout)
			case <-timer.C:
				// 这里已经超时了，再去检查一下还有没有任务，避免竞争的一种极端情况
				// 即协程池中只有一个worker，在这个worker快超时时，恰好来了一个任务，
				// 这样select中两个case都可执行，随机选择到了超时这个case，那么最后一个worker退出
				// 此时调用了close，就会导致没有worker运行这个任务，就不能保证所有任务都执行完成
				select {
				case task, ok := <-p.taskChan:
					if !ok {
						return
					}
					task()
					timer.Reset(p.idleTimeout)
				default:
					return
				}
			}
		}
	}()
}

func (p *Pool) Close() {
	p.mu.Lock()
	if p.closed {
		p.mu.Unlock()
		return
	}
	p.closed = true
	// 如果队列里还有任务，但没有 worker 了，补一个收尾 worker~
	if len(p.taskChan) > 0 && p.workerCount == 0 {
		p.StartWorker()
	}
	close(p.taskChan)
	p.mu.Unlock()
	p.wg.Wait() // 等待所有协程结束
}

func main() {
	pool := NewPool(200, 8, 300*time.Millisecond)

	var submitted atomic.Int64
	var executed atomic.Int64

	runWave := func(waveName string, submitterCount int, tasksPerSubmitter int, sleepBetweenSubmit time.Duration) {
		fmt.Printf("\n========== %s ==========\n", waveName)

		var submitWG sync.WaitGroup
		submitWG.Add(submitterCount)

		start := time.Now()

		for i := 0; i < submitterCount; i++ {
			go func(submitterID int) {
				defer submitWG.Done()

				for j := 0; j < tasksPerSubmitter; j++ {
					taskID := fmt.Sprintf("%s-s%d-t%d", waveName, submitterID, j)

					err := pool.Submit(func() {
						// 模拟真实任务：有一点随机耗时
						time.Sleep(time.Duration(rand.Intn(5)+1) * time.Millisecond)

						// 模拟一点轻微计算
						x := 0
						for k := 0; k < 1000; k++ {
							x += k
						}
						_ = x

						n := executed.Add(1)
						if n%2000 == 0 {
							fmt.Printf("[exec] %s executed=%d\n", taskID, n)
						}
					})
					if err != nil {
						fmt.Printf("[submit err] %s err=%v\n", taskID, err)
						return
					}

					n := submitted.Add(1)
					if n%2000 == 0 {
						fmt.Printf("[submit] submitted=%d\n", n)
					}

					if sleepBetweenSubmit > 0 {
						time.Sleep(sleepBetweenSubmit)
					}
				}
			}(i)
		}

		// 额外打印一下池子状态
		done := make(chan struct{})
		go func() {
			ticker := time.NewTicker(200 * time.Millisecond)
			defer ticker.Stop()
			for {
				select {
				case <-ticker.C:
					pool.mu.Lock()
					workers := pool.workerCount
					queued := len(pool.taskChan)
					maxWorkers := pool.maxWorkerCount
					pool.mu.Unlock()

					fmt.Printf("[stat] workers=%d max=%d queued=%d submitted=%d executed=%d\n",
						workers, maxWorkers, queued, submitted.Load(), executed.Load())
				case <-done:
					return
				}
			}
		}()

		submitWG.Wait()
		close(done)

		fmt.Printf("%s submit done, cost=%v\n", waveName, time.Since(start))
	}

	// 第一波：大量并发提交
	runWave("wave-1", 20, 1000, 0)

	// 等一会儿，让任务处理得差不多
	time.Sleep(2 * time.Second)

	fmt.Println("\n---- 调大并行度到 12 ----")
	if err := pool.SetParallelism(12); err != nil {
		fmt.Println("set parallelism err:", err)
	}

	// 第二波：继续压任务
	runWave("wave-2", 10, 800, 0)

	// 故意空闲一段时间，让 worker 因 idleTimeout 自动退出
	fmt.Println("\n---- 故意空闲 2 秒，观察 worker 是否自动退出 ----")
	time.Sleep(2 * time.Second)

	pool.mu.Lock()
	fmt.Printf("[after idle] workers=%d queued=%d\n", pool.workerCount, len(pool.taskChan))
	pool.mu.Unlock()

	// 第三波：空闲后再来任务，看 worker 能不能重新拉起
	runWave("wave-3", 5, 500, 1*time.Millisecond)

	// 等任务自然跑完
	for {
		pool.mu.Lock()
		queued := len(pool.taskChan)
		workers := pool.workerCount
		pool.mu.Unlock()

		fmt.Printf("[drain] queued=%d workers=%d submitted=%d executed=%d\n",
			queued, workers, submitted.Load(), executed.Load())

		if queued == 0 {
			break
		}
		time.Sleep(300 * time.Millisecond)
	}

	fmt.Println("\n---- Close pool ----")
	pool.Close()

	finalSubmitted := submitted.Load()
	finalExecuted := executed.Load()

	fmt.Printf("\n========== RESULT ==========\n")
	fmt.Printf("submitted = %d\n", finalSubmitted)
	fmt.Printf("executed  = %d\n", finalExecuted)

	if finalSubmitted == finalExecuted {
		fmt.Println("PASS: 所有已提交任务都执行完成了")
	} else {
		fmt.Println("FAIL: 有任务丢失或未执行完")
	}
}
