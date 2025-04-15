package main

import (
	"fmt"
	"time"
)

func main() {
	// 协程 A
	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			fmt.Println("协程 A")
		}
	}()

	// 协程 B
	go func() {
		defer func() {
			if err := recover(); err != nil { // 哪个协程Panic就需要在哪个协程recover
				fmt.Println("协程 B recover:", err)
			}
		}()
		time.Sleep(1 * time.Second) // 确保 协程 A 先运行起来
		panic("协程 B panic")
	}()

	time.Sleep(10 * time.Second) // 充分等待协程 B 触发 panic 完成和协程 A 执行完毕
	fmt.Println("main end")
}
