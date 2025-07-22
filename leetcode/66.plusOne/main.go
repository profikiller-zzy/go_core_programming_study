package main

func plusOne(digits []int) []int {
	var (
		carry int
		cur   int
	)
	n := len(digits)
	digits[n-1]++
	for index := n - 1; index >= 0; index-- {
		cur = digits[index] + carry
		carry = cur / 10
		digits[index] = cur % 10
	}
	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}
	return digits
}
