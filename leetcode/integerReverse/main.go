package main

import (
	"fmt"
)

func reverse(x int) int {
	isNegative := false
	if x < 0 {
		x = -x
		isNegative = true
	}
	res := 0
	for x != 0 {
		low := x % 10
		res = res*10 + low
		x /= 10
	}
	if isNegative && res > 1<<31 || !isNegative && res >= 1<<31 {
		return 0
	}
	if isNegative {
		return -res
	}
	return res
}

func main() {
	fmt.Println(reverse(-123))
}
