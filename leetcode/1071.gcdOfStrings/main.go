package main

import "fmt"

func gcdOfStrings(str1 string, str2 string) string {
	len1, len2 := len(str1), len(str2)
	gcdLen := gcd(len1, len2)
	for l := gcdLen; l > 0; l-- {
		if gcdLen%l != 0 {
			continue
		}
		prefix := str1[0:l]
		l1, l2 := 0, 0
		for l1 < len1 || l2 < len2 {
			if l1 < len1 {
				if prefix != str1[l1:l1+l] {
					break
				}
				l1 += l
			}
			if l2 < len2 {
				if prefix != str2[l2:l2+l] {
					break
				}
				l2 += l
			}
		}
		if l1 == len1 && l2 == len2 {
			return prefix
		}
	}
	return ""
}

func gcd(m int, n int) int {
	for n != 0 {
		temp := m % n
		m = n
		n = temp
	}
	return m
}

func main() {
	str1 := "ABCABC"
	str2 := "ABC"
	fmt.Println(gcdOfStrings(str1, str2))
}
