package main

import (
	"fmt"
)

func main() {
	panicRecover(0)
}

func panicRecover(num int) {
	defer func() { // defer recover 需要定义在panic 之前
		if err := recover(); err != nil {
			fmt.Println("捕获到 panic:", err)
		}
	}()
	if num == 0 {
		panic("发生了 panic") // panic之后，函数退出，开始沿调用栈向上寻找 defer。
	}
	fmt.Println("num:", num)
}
