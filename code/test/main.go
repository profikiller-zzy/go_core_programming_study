package main

import (
	"fmt"
	"time"
)

func loopUnrolling_1(count int) int {
	var sum int
	for i := 0; i < count; i++ {
		sum += i
	}
	return sum
}

func loopUnrolling_2(count int) int {
	var sum int
	for i := 0; i < count; i += 4 {
		sum += i
		sum += i + 1
		sum += i + 2
		sum += i + 3
	}
	return sum
}

func loopUnrolling_3(count int) int {
	var sum1, sum2, sum3, sum4 int
	for i := 0; i < count; i += 4 {
		sum1 += i
		sum2 += i + 1
		sum3 += i + 2
		sum4 += i + 3
	}
	return sum1 + sum2 + sum3 + sum4
}

// 测试函数
func benchmarkFunc(name string, f func(count int) int, count int) {
	start := time.Now()          // 获取当前时间
	result := f(count)           // 调用待测函数
	elapsed := time.Since(start) // 计算函数执行时间
	fmt.Printf("%s: result=%d, elapsed=%s\n", name, result, elapsed)
}

func main() {
	count := 100000000 // 设置测试的循环次数

	// 测试 func loopUnrolling_1
	benchmarkFunc("loopUnrolling_1", loopUnrolling_1, count)

	// 测试 func loopUnrolling_2
	benchmarkFunc("loopUnrolling_2", loopUnrolling_2, count)

	// 测试 func loopUnrolling_3
	benchmarkFunc("loopUnrolling_3", loopUnrolling_3, count)
}
