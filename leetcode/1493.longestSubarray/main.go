package main

import "fmt"

func longestSubarray(nums []int) int {
	// 等价于寻找最多包含一个0的子序列长度再减1
	var left, right int
	zeroCount := 1
	for right < len(nums) {
		if nums[right] == 0 {
			zeroCount--
		}
		right++

		if zeroCount < 0 {
			if nums[left] == 0 {
				zeroCount++
			}
			left++
		}
	}
	return right - left - 1
}

func main() {
	nums := []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	fmt.Println(longestSubarray(nums))
}
