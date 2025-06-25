package main

import (
	"fmt"
	"time"
)

func spawnGoroutine() {
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine done")
	}()
	fmt.Println("spawnGoroutine returned")
}

func main() {
	spawnGoroutine()
	time.Sleep(3 * time.Second)
	fmt.Println("main done")
}
