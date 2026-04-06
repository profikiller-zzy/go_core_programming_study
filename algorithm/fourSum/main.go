package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	var res [][]int
	n := len(nums)

	sort.Ints(nums) // 先对数组进行排序
	for i := 0; i < n-3; i++ {
		// 跳过相等的重复值
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j+1, n-1

			// 使用双指针查找剩余的两个数
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					// 分别跳过右边和左边重复的值
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					right--
					left++
				} else if sum > target { // 可能是右边的数字太大了
					right--
				} else {
					left++
				}
			}
		}
	}
	return res
}

func main() {
	var nums []int
	nums = []int{2, 2, 2, 2, 2}
	fmt.Println(fourSum(nums, 8))
}
