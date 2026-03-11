package main

import "fmt"

func pivotIndex(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	leftSum := 0
	for i, num := range nums {
		rightSum := total - leftSum - num
		if leftSum == rightSum {
			return i
		}
		leftSum += num
	}

	return -1
}

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(nums))
}
