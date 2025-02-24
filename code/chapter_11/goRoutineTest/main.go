package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go printMessage("Hello from Goroutine") // 启动一个协程
	printMessage("Hello from Main Thread")  // 主线程执行
}
