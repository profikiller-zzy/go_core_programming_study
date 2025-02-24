package main

import (
	"fmt"
	"time"
)

// 需求：现在要计算 1-200 的各个数的阶乘，并且把各个数的阶乘放入到 map 中。
// 最后显示出来。要求使用 goroutine 完成
// 思路
// 1. 编写一个函数，来计算各个数的阶乘，并放入到 map 中.
// 2. 我们启动的协程多个，统计的将结果放入到 map 中
// 3. map 应该做出一个全局的.

var (
	myMap = make(map[int]int, 10)
)

// test 函数就是计算 n!, 让将这个结果放入到 myMap
func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//这里我们将 res 放入到 myMap
	myMap[n] = res //concurrent map writes?
}

func main() {
	// 我们这里开启多个协程完成这个任务[200 个]
	for i := 1; i <= 200; i++ {
		go test(i)
	}
	//休眠 10 秒钟【第二个问题 】
	time.Sleep(time.Second * 10)
	//这里我们输出结果,变量这个结果
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
}
