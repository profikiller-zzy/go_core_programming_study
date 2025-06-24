package main

// 使用go语言实现漏桶算法
import (
	"fmt"
	"sync"
	"time"
)

// LeakyBucket 限流器
type LeakyBucket struct {
	capacity int           // 桶的容量
	interval time.Duration // 漏水间隔
	bucket   chan struct{}
	stopChan chan struct{}
	stopOnce sync.Once
}

// NewLeakyBucket 创建一个新的 LeakyBucket
func NewLeakyBucket(capacity int, leakInterval time.Duration) *LeakyBucket {
	lb := &LeakyBucket{
		capacity: capacity,
		interval: leakInterval,
		bucket:   make(chan struct{}, capacity),
		stopChan: make(chan struct{}),
	}

	// 启动漏水协程
	go lb.leak()
	return lb
}

// leak 模拟漏桶定期漏水
func (lb *LeakyBucket) leak() {
	ticker := time.NewTicker(lb.interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			select {
			case <-lb.bucket:
				// 成功漏出一滴水
			default:
				// 桶是空的，什么都不做
			}
		case <-lb.stopChan:
			// 这里stopChan被关闭后，再从它那里接收是不会阻塞的，并且假设channel中没有值，则接收会返回一个零值
			// 优雅退出后台
			return
		}
	}
}

// Allow 判断是否允许请求进入桶中
func (lb *LeakyBucket) Allow() bool {
	select {
	case lb.bucket <- struct{}{}:
		return true
	default:
		// 桶满了
		return false
	}
}

// Stop 停止漏水
func (lb *LeakyBucket) Stop() {
	lb.stopOnce.Do(func() { // 确保只被调用一次，防止多次调用close channel导致程序崩溃
		close(lb.stopChan)
	})
}

func main() {
	// 创建一个容量为 5，100ms 漏一滴的漏桶
	lb := NewLeakyBucket(5, 500*time.Millisecond)
	defer lb.Stop()

	// 模拟一秒钟内每 50ms 发一个请求
	ticker := time.NewTicker(250 * time.Millisecond)
	defer ticker.Stop()

	// 创建一个timer，一秒之后Timer到期，当时的时间会被发送给Timer中的channel
	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()

	for {
		select {
		case <-ticker.C:
			if lb.Allow() {
				fmt.Println(time.Now().Format("15:04:05.000"), "Request allowed ✅")
			} else {
				fmt.Println(time.Now().Format("15:04:05.000"), "Request denied ❌")
			}
		case <-timer.C:
			fmt.Println("测试结束")
			return
		}
	}
}
