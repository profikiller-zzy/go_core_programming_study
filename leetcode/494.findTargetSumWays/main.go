package main

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	total := sum + target
	if total < 0 || total%2 != 0 {
		return 0
	}

	capacity := total / 2
	dp := make([]int, capacity+1)
	dp[0] = 1

	for _, num := range nums {
		// 倒序：每个 num 只能选一次
		for j := capacity; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}

	return dp[capacity]
}
