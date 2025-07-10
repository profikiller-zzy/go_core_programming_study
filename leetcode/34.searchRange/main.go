package main

import "fmt"

// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRange(nums []int, target int) []int {
	var res []int = make([]int, 2)
	res[0] = -1
	res[1] = -1
	if len(nums) == 0 {
		return res
	}
	left, right, mid := 0, len(nums)-1, 0
	// 不断压缩右边界，找到最左边的边界
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	// 这时候left一般是指向范围的最左侧边界，如果left越界或者left指向的num不等于target，说明没有找到target
	if left <= len(nums)-1 && nums[left] == target {
		res[0] = left
	} else {
		res[0] = -1
	}

	// 第二次二分循环结束时，left指向的是第一个大于target的元素位置（或者是越界），所right是最后一个大雨等于target的元素位置
	left, right = 0, len(nums)-1
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if right >= 0 && nums[right] == target {
		res[1] = right
	} else {
		res[1] = -1
	}
	return res
}

func main() {
	var nums []int = []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(nums, 7))
}
