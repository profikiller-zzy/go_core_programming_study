package main

import "strings"

func removeKdigits(num string, k int) string {
	if k == 0 {
		return num
	}
	if len(num) <= k {
		return "0"
	}

	// 预分配栈空间，直接使用字节操作
	stack := make([]byte, 0, len(num))
	for i := 0; i < len(num); i++ {
		ch := num[i]
		// 当栈非空、仍有删除额度且当前字符小于栈顶，则弹出栈顶
		for len(stack) > 0 && k > 0 && ch < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		if k == 0 {
			stack = append(stack, num[i:]...)
			break
		}
		stack = append(stack, ch)
	}

	if k > 0 {
		stack = stack[:len(stack)-k]
	}
	result := strings.TrimLeft(string(stack), "0")
	if result == "" {
		return "0"
	}
	return result
}
