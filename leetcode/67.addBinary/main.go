package main

import "strings"

func reverse(str string) string {
	var res strings.Builder
	for i := len(str) - 1; i >= 0; i-- {
		res.WriteByte(str[i])
	}
	return res.String()
}

func addBinary(a string, b string) string {
	var res strings.Builder
	var carry, cur byte
	for index1, index2 := len(a)-1, len(b)-1; index1 >= 0 || index2 >= 0 || carry > 0; index1, index2 = index1-1, index2-1 {
		cur = carry
		if index1 >= 0 {
			cur += a[index1] - '0'
		}
		if index2 >= 0 {
			cur += b[index2] - '0'
		}
		res.WriteByte(cur%2 + '0')
		carry = cur / 2
	}
	return reverse(res.String())
}
