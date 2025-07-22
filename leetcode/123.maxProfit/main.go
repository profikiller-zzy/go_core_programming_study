package main

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	buy1, sell1, buy2, sell2 := -prices[0], 0, -prices[0], 0
	for index := 1; index < len(prices); index++ {
		buy1 = max(buy1, -prices[index])
		sell1 = max(sell1, buy1+prices[index])
		buy2 = max(buy2, sell1-prices[index])
		sell2 = max(sell2, buy2+prices[index])
	}
	return sell2
}
