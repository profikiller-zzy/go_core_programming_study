package main

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)

	// 状态转移方程 dp[i][j] 表示 text1[0:i] 和 text2[0:j] 的最长公共子序列长度
	var dp = make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	// 初始化 dp 数组
	for i := 0; i < m+1; i++ {
		dp[i][0] = 0
	}
	for j := 0; j < n+1; j++ {
		dp[0][j] = 0
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

func main() {

}
