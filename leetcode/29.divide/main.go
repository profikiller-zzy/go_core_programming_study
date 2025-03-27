package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 {
		if divisor == 1 {
			return math.MinInt32
		}
		if divisor == -1 {
			return math.MaxInt32
		}
	}
	if divisor == math.MinInt32 {
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}

	negative := (dividend ^ divisor) < 0 // 判断符号是否不同
	// 将被除数和除数转为正数处理
	dividendAbs, divisorAbs := abs(dividend), abs(divisor)
	quotient := 0

	for dividendAbs >= divisorAbs {
		currentDivisor, currentStep := divisorAbs, 1
		// 找到最大的 currentDivisor = divisor * 2^N，满足 currentDivisor <= dividendAbs
		for currentDivisor <= (dividendAbs >> 1) { // 判断条件是currentDivisor小于等于当前除数的一半，防止翻倍后溢出
			currentDivisor <<= 1 // 翻倍
			currentStep <<= 1    // 翻倍
		}
		quotient += currentStep
		dividendAbs -= currentDivisor
	}

	if negative {
		quotient = -quotient
	}
	// 处理溢出
	if quotient > math.MaxInt32 {
		return math.MaxInt32
	}
	if quotient < math.MinInt32 {
		return math.MinInt32
	}
	return quotient
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(divide(1, 1))
}
