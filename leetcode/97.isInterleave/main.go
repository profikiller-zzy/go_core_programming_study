package main

import (
	"fmt"
)

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	// 这里dp[i][j] 表示s1[0:i-1] 和 s2[0:j-1] 能否交错组成 s3[0:i+j-1] (i, j 大于零)
	dp := make([][]bool, len(s1)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s2)+1)
	}
	// 初始化 dp 数组
	dp[0][0] = true
	for i := 1; i <= len(s1); i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}
	for j := 1; j <= len(s2); j++ {
		dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if i > 0 && s1[i-1] == s3[i+j-1] {
				dp[i][j] = dp[i][j] || dp[i-1][j]
			}
			if j > 0 && s2[j-1] == s3[i+j-1] {
				dp[i][j] = dp[i][j] || dp[i][j-1]
			}
		}
	}
	return dp[len(s1)][len(s2)]
}

func main() {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"
	fmt.Println(isInterleave(s1, s2, s3))
}
