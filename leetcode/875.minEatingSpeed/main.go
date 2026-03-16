package main

import "fmt"

func minEatingSpeed(piles []int, h int) int {
	var maxPile int
	for _, pile := range piles {
		if pile > maxPile {
			maxPile = pile
		}
	}
	calTime := func(k int) int {
		var t int
		for _, pile := range piles {
			t += (pile-1)/k + 1
		}
		return t
	}
	left, right := 1, maxPile
	for left <= right {
		mid := (left + right) / 2
		costT := calTime(mid)
		if costT > h {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func main() {
	piles := []int{30, 11, 23, 4, 20}
	fmt.Println(minEatingSpeed(piles, 5))
}
