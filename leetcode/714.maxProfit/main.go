package main

import "fmt"

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	buy, sell := make([]int, n), make([]int, n)
	buy[0] = -prices[0] // 第i天交易完之后手里有股票
	sell[0] = 0         // 第i天交易完之后手里没有股票
	for i := 1; i < n; i++ {
		buy[i] = max(sell[i-1]-prices[i], buy[i-1])
		sell[i] = max(buy[i]+prices[i]-fee, sell[i-1])
	}
	return sell[n-1]
}

func main() {
	prices := []int{1, 3, 7, 5, 10, 3}
	fee := 3
	fmt.Println(maxProfit(prices, fee))
}
