package main

import "strings"

func isPalindrome(s string) bool {
	var processed strings.Builder
	for index := 0; index < len(s); index++ {
		if s[index] >= 'a' && s[index] <= 'z' || s[index] >= 'A' && s[index] <= 'Z' || s[index] >= '0' && s[index] <= '9' {
			if s[index] >= 'A' && s[index] <= 'Z' {
				processed.WriteByte(s[index] + 32) // 转为小写
			} else {
				processed.WriteByte(s[index])
			}
		} else {
			continue // 忽略非字母和数字
		}
	}
	processedStr := processed.String()
	for left, right := 0, len(processedStr)-1; left < right; {
		if processedStr[left] != processedStr[right] {
			return false
		}
		left++
		right--
	}
	return true
}
