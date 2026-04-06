package main

import (
	"fmt"
	"sync"
	"time"
)

// UnlimitedChan 无限制缓存channel
type UnlimitedChan struct {
	mu     sync.Mutex    // 保护数据访问
	cond   *sync.Cond    // 用于等待/通知
	buffer []interface{} // 存储数据的缓冲区
	closed bool          // 标记channel是否关闭
}

// NewUnlimitedChan 创建无限制缓存channel
func NewUnlimitedChan() *UnlimitedChan {
	uc := &UnlimitedChan{}
	uc.cond = sync.NewCond(&uc.mu) // 绑定互斥锁
	return uc
}

// Send 发送数据到channel
func (uc *UnlimitedChan) Send(data interface{}) {
	uc.mu.Lock()
	defer uc.mu.Unlock()

	if uc.closed {
		panic("send on closed UnlimitedChan")
	}

	uc.buffer = append(uc.buffer, data)
	uc.cond.Signal() // 唤醒一个等待的接收者
}

// Recv 从channel接收数据
func (uc *UnlimitedChan) Recv() (interface{}, bool) {
	uc.mu.Lock()
	defer uc.mu.Unlock()

	// 若缓冲区为空且未关闭，等待数据
	for len(uc.buffer) == 0 && !uc.closed {
		uc.cond.Wait() // 释放锁并等待，被唤醒时重新获取锁
	}

	// 若已关闭且缓冲区为空，返回零值和false
	if len(uc.buffer) == 0 && uc.closed {
		return nil, false
	}

	// 取出缓冲区第一个元素
	data := uc.buffer[0]
	uc.buffer = uc.buffer[1:] // 移除已接收元素
	return data, true
}

// Close 关闭channel
func (uc *UnlimitedChan) Close() {
	uc.mu.Lock()
	defer uc.mu.Unlock()

	if uc.closed {
		return // 避免重复关闭
	}

	uc.closed = true
	uc.cond.Broadcast() // 唤醒所有等待的接收者
}

// Chan 返回可迭代的channel
func (uc *UnlimitedChan) Chan() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		for {
			data, ok := uc.Recv()
			if !ok {
				close(ch)
				return
			}
			ch <- data
		}
	}()
	return ch
}

func main() {
	c := make(chan int, 1)
	go func() {
		for i := 0; i < 3; i++ {
			c <- i
		}
	}()
	x := <-c
	fmt.Println(x)
	close(c)
	time.Sleep(time.Second)
}
