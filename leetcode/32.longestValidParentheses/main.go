package main

//https://leetcode.cn/problems/longest-valid-parentheses/

func longestValidParentheses(s string) int {
	var ans int
	var stack []int
	stack = append(stack, -1) // 初始化栈，栈底元素为-1，防止直接遇到右括号导致栈为空
	for i := 0; i < len(s); i++ {
		if s[i] == '(' { // 每次只需要记录左括号的位置
			stack = append(stack, i)
		} else { // 遇到右括号，弹出栈顶元素，也就是匹配到的左括号弹出
			stack = stack[:len(stack)-1]
			if len(stack) == 0 { // 栈为空，说明没有匹配的左括号，重新从当前位置开始寻找
				stack = append(stack, i)
			} else { // 栈不为空，计算当前合法括号的长度
				ans = max(ans, i-stack[len(stack)-1])
			}
		}
	}
	return ans
}

func main() {

}
