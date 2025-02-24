package main

import (
	"fmt"
	"golang.org/x/net/context"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // 监听取消信号
			fmt.Println("Worker canceled:", ctx.Err())
			return
		default:
			fmt.Println("Working...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)

	time.Sleep(3 * time.Second)
	cancel()                    // 手动取消 worker
	time.Sleep(1 * time.Second) // 等待 worker 退出
}
