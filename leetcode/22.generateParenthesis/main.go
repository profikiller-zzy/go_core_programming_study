package main

func generateParenthesis(n int) []string {
	var path []byte
	var res []string
	var left, right int
	var dfs func()
	dfs = func() {
		if right == n {
			res = append(res, string(path))
			return
		}
		if left < n {
			path = append(path, '(')
			left++
			dfs()
			left--
			path = path[:len(path)-1]
		}
		if right < left {
			path = append(path, ')')
			right++
			dfs()
			right--
			path = path[:len(path)-1]
		}
	}
	dfs()
	return res
}
