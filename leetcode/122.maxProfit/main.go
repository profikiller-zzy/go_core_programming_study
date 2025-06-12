package main

func maxProfit(prices []int) int {
	var m int
	for index := 1; index < len(prices); index++ {
		if prices[index] > prices[index-1] {
			m += prices[index] - prices[index-1]
		}
	}
	return m
}
