package main

import "strconv"

func isPalindrome(x int) bool {
	xStr := strconv.Itoa(x)
	left, right := 0, len(xStr)-1
	for left < right {
		if xStr[left] != xStr[right] {
			return false
		}
		left++
		right--
	}
	return true
}
