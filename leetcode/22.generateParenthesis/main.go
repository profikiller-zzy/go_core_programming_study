package main

func generateParenthesis(n int) []string {
	var result []string
	var backtrack func(string, int, int)
	backtrack = func(cur string, left int, right int) {
		if len(cur) == 2*n {
			result = append(result, cur)
			return
		}
		if left < n {
			backtrack(cur+"(", left+1, right) // 只有当左括号数量 left < n 时才允许添加新的左括号
		}
		if right < left {
			backtrack(cur+")", left, right+1) // 只有当右括号数量 right < left 时才允许添加右括号
		}
	}
	backtrack("", 0, 0)
	return result
}
