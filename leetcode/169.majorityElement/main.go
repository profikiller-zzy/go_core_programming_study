package main

import "fmt"

func majorityElement(nums []int) int {
	var (
		curNum, count int
	)
	for _, num := range nums {
		if count == 0 {
			curNum = num
			count++
		} else if curNum == num {
			count++
		} else {
			count--
		}
	}
	return curNum
}

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
}
