package main

import "fmt"

//https://leetcode.cn/problems/search-in-rotated-sorted-array/description/

func search1(nums []int, target int) int {
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

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	// 第一步：找到旋转点 k (最小值的下标)
	low, high := 0, n-1
	for low < high {
		mid := (low + high) / 2
		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	k := low

	// 第二步：在逻辑上的“有序数组”中进行二分查找
	// 逻辑范围是 [0, n-1]，但实际访问时需要通过 (mid + k) % n 映射回原数组
	low, high = 0, n-1
	for low <= high {
		mid := (low + high) / 2
		realMid := (mid + k) % n
		if nums[realMid] == target {
			return realMid
		} else if nums[realMid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	var nums = []int{1, 2, 4, 5, 6, 7, 0}
	fmt.Println(search(nums, 0))
}
