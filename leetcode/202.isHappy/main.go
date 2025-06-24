package main

func isHappy(n int) bool {
	slow, fast := n, bitSquareSum(n)
	for slow != fast {
		slow = bitSquareSum(slow)
		fast = bitSquareSum(fast)
		fast = bitSquareSum(fast)
	}
	return slow == 1
}

func bitSquareSum(src int) int {
	var sum int
	for src != 0 {
		bit := src % 10
		sum += bit * bit
		src = src / 10
	}
	return sum
}
