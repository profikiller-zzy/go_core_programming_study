package main

type StockSpanner struct {
	minStack []int
	prices   []int
}

func Constructor() StockSpanner {
	return StockSpanner{
		minStack: make([]int, 0),
		prices:   make([]int, 0),
	}
}

func (this *StockSpanner) Next(price int) int {
	for len(this.minStack) > 0 && price >= this.prices[this.minStack[len(this.minStack)-1]] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
	this.prices = append(this.prices, price)
	var res int
	if len(this.minStack) == 0 {
		res = len(this.prices)
	} else {
		res = len(this.prices) - 1 - this.minStack[len(this.minStack)-1]
	}
	this.minStack = append(this.minStack, len(this.prices)-1)
	return res
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
