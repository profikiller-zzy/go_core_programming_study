package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left, right, mid := 0, len(nums)-1, 0
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	var nums = []int{1, 3, 5, 6}
	target := 7
	fmt.Println(searchInsert(nums, target))
}
