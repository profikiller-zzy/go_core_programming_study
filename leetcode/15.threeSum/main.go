package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/3sum

func threeSum(nums []int) [][]int {
	var res [][]int
	if len(nums) < 3 {
		return res
	}
	// 题目提示了输出的顺序和三元组的顺序不重要，于是先将nums排序
	sort.Ints(nums)
	for left := 0; left < len(nums)-2; left++ {
		if nums[left] > 0 { // 假如三元组里最小的数都大于零，那么以及没有三元组的和等于零了
			break
		}
		if left > 0 && nums[left] == nums[left-1] { // 跳过重复的
			continue
		}
		mid, right := left+1, len(nums)-1
		for mid < right {
			sum := nums[left] + nums[mid] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[left], nums[mid], nums[right]})
				for mid < right && nums[mid] == nums[mid+1] { // 跳过重复的
					mid++
				}
				for mid < right && nums[right] == nums[right-1] { // 跳过重复的
					right--
				}
				mid++
				right--
			} else if sum > 0 { // 这种情况可能是右边的数太大了
				right--
			} else {
				mid++
			}
		}
	}
	return res
}

func main() {
	var nums []int
	nums = []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
