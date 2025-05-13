package main

import "fmt"

//https://leetcode.cn/problems/longest-valid-parentheses/

func longestValidParentheses(s string) int {
	var ans int
	// 栈底元素为当前已经遍历过的元素中「最后一个没有被匹配的右括号的下标」，初始化栈底元素为-1
	var stack []int
	stack = append(stack, -1)
	for index := 0; index < len(s); index++ {
		if s[index] == '(' { // 每次只需要记录左括号的位置
			stack = append(stack, index)
		} else { // 遇到右括号，弹出栈顶元素，也就是匹配到的左括号弹出
			stack = stack[:len(stack)-1]
			if len(stack) == 0 { // 栈为空，说明没有匹配的左括号，重新从当前位置开始寻找
				stack = append(stack, index)
			} else { // 栈不为空，计算当前合法括号的长度
				ans = max(ans, index-stack[len(stack)-1])
			}
		}
	}
	return ans
}

func main() {
	s := "))()())"
	fmt.Println(longestValidParentheses(s)) // 输出 4
}
