package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	// 先初始化dp数组
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i < len(dp); i++ {
		dp[i][0] = i
	}
	for j := 0; j < len(dp[0]); j++ {
		dp[0][j] = j
	}
	for i := 0; i < len(word1); i++ {
		for j := 0; j < len(word2); j++ {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = min(dp[i][j], dp[i+1][j], dp[i][j+1]) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func main() {
	fmt.Println(minDistance("leetcode", "code"))
	fmt.Println(minDistance("horse", "ros"))
}
