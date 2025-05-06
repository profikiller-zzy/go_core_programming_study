package main

import "fmt"

//https://leetcode.cn/problems/search-in-rotated-sorted-array/description/

func search(nums []int, target int) int {
	// 先找到k
	var (
		low  = 0
		high = len(nums) - 1
		k    = 0
	)
	for low < high {
		mid := (low + high) / 2
		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	k = low // low == high，最终指向最小值下标

	// 然后根据k进行二分查找
	var binarySearch func(nums []int, target int, low int, high int) int
	binarySearch = func(nums []int, target int, low int, high int) int {
		mid := (low + high) / 2
		for low <= high {
			if nums[mid] == target {
				return mid
			} else if nums[mid] < target {
				low = mid + 1
			} else {
				high = mid - 1
			}
			mid = (low + high) / 2
		}
		return -1
	}
	if k == 0 {
		return binarySearch(nums, target, 0, len(nums)-1)
	} else {
		if target >= nums[0] {
			return binarySearch(nums, target, 0, k-1)
		} else {
			return binarySearch(nums, target, k, len(nums)-1)
		}
	}
}

func main() {
	var nums = []int{1, 2, 4, 5, 6, 7, 0}
	fmt.Println(search(nums, 0))
}
