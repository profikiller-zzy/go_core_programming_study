package main

func maxProfit1(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	var dp = make([]int, len(prices))
	var maxP int
	dp[0] = 0
	dp[1] = prices[1] - prices[0]
	maxP = dp[1]
	for i := 2; i < len(prices); i++ {
		dp[i] = max(prices[i]-prices[i-1], dp[i-1]+prices[i]-prices[i-1])
		maxP = max(maxP, dp[i])
	}
	if maxP < 0 {
		return 0
	}
	return maxP
}

func maxProfit(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}
	buy, sell := -prices[0], 0
	for index := 1; index < len(prices); index++ {
		buy = max(buy, -prices[index])
		sell = max(sell, buy+prices[index])
	}
	return sell
}
