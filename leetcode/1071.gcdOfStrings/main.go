package main

import "fmt"

func gcdOfStrings(str1 string, str2 string) string {
	l1, l2 := len(str1), len(str2)
	factor := gcd(l1, l2)
	prefix := str1[:factor]
	var cur1, cur2 int
	for cur1 < len(str1) || cur2 < len(str2) {
		if cur1 < len(str1) {
			if str1[cur1:cur1+factor] != prefix {
				return ""
			}
			cur1 += factor
		}
		if cur2 < len(str2) {
			if str2[cur2:cur2+factor] != prefix {
				return ""
			}
			cur2 += factor
		}
	}
	return prefix
}

func gcd(m, n int) int {
	for n != 0 {
		factor := m % n
		m = n
		n = factor
	}
	return m
}

func main() {
	str1 := "ABCABC"
	str2 := "ABC"
	fmt.Println(gcdOfStrings(str1, str2))
}
