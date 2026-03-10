package main

func longestPalindrome1(s string) string {
	if len(s) <= 1 {
		return s
	}
	n := len(s)
	var res string
	res = s[0:1]

	// dp[left][j] 表示下标从 left 到 j 的子串是否是回文串
	var dp = make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	// 初始化
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	for left := n - 1; left >= 0; left-- {
		for right := left + 1; right <= n-1; right++ {
			if right-left == 1 { // 当前这个子串的长度为2
				dp[left][right] = s[left] == s[right]
			} else {
				dp[left][right] = s[left] == s[right] && dp[left+1][right-1]
			}
			if dp[left][right] {
				if len(res) < right-left+1 {
					res = s[left : right+1] // 更新结果
				}
			}
		}
	}
	return res
}

// longestPalindrome 中心扩展法解最长回文串
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	expandAroundCenter := func(left, right int) (int, int) {
		for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
		}
		return left + 1, right - 1
	}
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(i, i)
		left2, right2 := expandAroundCenter(i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func main() {
	s := "babad"
	println(longestPalindrome(s)) // 输出 "bab" 或 "aba"
}
