package main

import "fmt"

func isMatch(s string, p string) bool {
	m := len(s)
	n := len(p)
	var dp [][]bool = make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true

	// 处理空字符串与模式 p 的匹配，处理类似 "a*", ".*" 这种匹配空字符串的情况
	for j := 2; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	// 使用状态转移方程来填充dp二维数组
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '.' || p[j-1] == s[i-1] {
				// 如果当前模式串字符是'.'或者是当前模式串字符和目标串字符匹配
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				// '*' 匹配零个或者多个模式串前一个字符
				dp[i][j] = dp[i][j-2] // 这种情况'*'匹配零个字符
				if p[j-2] == '.' || p[j-2] == s[i-1] {
					// '*' 匹配一个或多个前面的字符
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}
		}
	}

	return dp[m][n]
}

func main() {
	s := "aab"
	p := "c*a*b"
	fmt.Println(isMatch(s, p))
}
