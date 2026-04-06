package main

import (
	"fmt"
	"sort"
)

func coinChange1(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = i
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				// 核心状态转移方程：
				// dp[i] = min(当前值, 减去硬币后的值 + 1)
				if dp[i-coin]+1 < dp[i] {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	sort.Ints(coins)
	queue := []int{amount}
	step := 0
	visited := make([]bool, amount+1)
	visited[amount] = true
	for len(queue) > 0 {
		size := len(queue)
		step++
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			for j := 0; j < len(coins); j++ {
				next := cur - coins[j]
				if next == 0 {
					return step
				} else if next > 0 && !visited[next] {
					visited[next] = true
					queue = append(queue, next)
				}
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(coinChange([]int{2}, 3))
}
