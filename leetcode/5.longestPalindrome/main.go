package main

func longestPalindrome1(s string) string {
	if len(s) <= 1 {
		return s
	}
	n := len(s)
	var res string
	res = s[0:1]

	// dp[i][j] 表示下标从 i 到 j 的子串是否是回文串
	var dp = make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	// 初始化
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j <= n-1; j++ {
			if j-i == 1 { // 当前这个子串的长度为2
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
			}
			if dp[i][j] {
				if len(res) < j-i+1 {
					res = s[i : j+1] // 更新结果
				}
			}
		}
	}
	return res
}

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}

func main() {
	s := "babad"
	println(longestPalindrome(s)) // 输出 "bab" 或 "aba"
}
