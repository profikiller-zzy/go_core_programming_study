package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	res := 0
	isNegative := false
	state := 0
	maxInt := math.MaxInt32
	minInt := math.MinInt32

	for _, ch := range s {
		switch state {
		case 0:
			{
				if ch == ' ' {
					continue
				} else if ch == '+' {
					state = 1
				} else if ch == '-' {
					state = 1
					isNegative = true
				} else if ch >= '0' && ch <= '9' {
					res = res*10 + int(ch) - 48
					state = 2
				} else {
					state = 3 // end
				}
			}
		case 1:
			{
				if ch >= '0' && ch <= '9' {
					res = res*10 + int(ch) - 48
					state = 2
				} else {
					state = 3
				}
			}
		case 2:
			{
				if ch >= '0' && ch <= '9' {
					digit := int(ch) - 48

					// 检查是否即将超出范围
					if isNegative {
						if -res < (minInt+digit)/10 {
							return minInt
						}
					} else {
						if res > (maxInt-digit)/10 {
							return maxInt
						}
					}

					res = res*10 + digit
				} else {
					state = 3
				}
			}
		case 3:
			{
				break
			}
		}
	}

	if isNegative {
		res = -res
	}
	return res
}

func main() {
	res := myAtoi("9223372036854775808")
	fmt.Println(res)
}
