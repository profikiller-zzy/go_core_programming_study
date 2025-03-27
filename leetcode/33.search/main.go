package main

import "fmt"

//https://leetcode.cn/problems/search-in-rotated-sorted-array/description/

func search(nums []int, target int) int {
	// 先找到k
	var (
		low  int = 0
		high int = len(nums) - 1
	)
	kMid := (low + high) / 2
	k := 0
	for kMid >= 0 && kMid < len(nums) {
		kMid = (low + high) / 2
		if nums[kMid] > nums[high] { // 说明k在mid的右边，缩短
			low = kMid + 1
		} else if nums[kMid] < nums[high] { // 说明k在mid的左边，缩短
			high = kMid
		} else {
			k = kMid
			break
		}
	}

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
