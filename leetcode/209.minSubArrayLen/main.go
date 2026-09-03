package main

import (
	"math"
	"sort"
)

// minSubArrayLen1 前缀和+二分
func minSubArrayLen1(target int, nums []int) int {
	preSum := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	result := math.MaxInt32
	for i := 1; i <= len(nums); i++ {
		t := preSum[i-1] + target
		j := sort.Search(len(preSum), func(j int) bool {
			return preSum[j] >= t
		})
		if j < len(preSum) {
			result = min(result, j-i+1)
		}
	}
	if result == math.MaxInt32 {
		return 0
	}
	return result
}

// minSubArrayLen 滑动窗口
func minSubArrayLen(target int, nums []int) int {
	var start, end, curSum int
	result := math.MaxInt32
	for end < len(nums) {
		curSum += nums[end]
		end++
		for curSum >= target {
			result = min(result, end-start)
			curSum -= nums[start]
			start++
		}
	}
	if result == math.MaxInt32 {
		return 0
	}
	return result
}
