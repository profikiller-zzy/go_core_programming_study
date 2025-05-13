package main

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var dp = make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for index := 2; index < len(nums); index++ {
		dp[index] = max(nums[index]+dp[index-2], dp[index-1])
	}
	return dp[len(nums)-1]
}
