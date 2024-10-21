package main

import "fmt"

func maxNum(nums ...int) int {
	m := nums[0]
	for num := range nums {
		if num > m {
			m = num
		}
	}
	return m
}

func main() {
	fmt.Println(-10 % 3)
	fmt.Println(10 % -3)
	fmt.Println(maxNum(1, 2, 4, 5, 3, 0))
}
