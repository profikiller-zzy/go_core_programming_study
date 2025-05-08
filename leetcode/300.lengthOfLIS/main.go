package main

func lengthOfLIS(nums []int) int {
	var dp = make([]int, len(nums))
	dp[0] = 1
	maxLen := 1
	for index := 1; index < len(nums); index++ {
		var maxCurLen = 1
		for j := 0; j < index; j++ {
			if nums[index] > nums[j] {
				if maxCurLen < dp[j]+1 {
					maxCurLen = dp[j] + 1
				}
			}
		}
		dp[index] = maxCurLen
		if maxLen < maxCurLen {
			maxLen = maxCurLen
		}
	}
	return maxLen
}
