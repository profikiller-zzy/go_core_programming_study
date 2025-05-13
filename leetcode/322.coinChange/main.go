package main

import "sort"

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	sort.Ints(coins) // 先讲 coins 排序
	var dp = make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		// 求 dp[i]
		if i < coins[0] { // 如果 i 小于 coins[0]，说明没有 coins 可以使用
			dp[i] = -1
			continue
		}
		var minSteps = -1
		for j := 0; j < len(coins); j++ {
			if coins[j] > i { // 如果 coins[j] 大于 i，之后的 coins[j] 都大于 i，都不可用
				break
			}
			if dp[i-coins[j]] == -1 { // 如果 dp[i-coins[j]] == -1，说明没有 coins 可以使用
				continue
			}
			if minSteps == -1 {
				minSteps = dp[i-coins[j]]
			} else {
				minSteps = min(minSteps, dp[i-coins[j]])
			}
		}
		if minSteps == -1 {
			dp[i] = -1
		} else {
			dp[i] = minSteps + 1
		}
	}
	return dp[amount]
}
