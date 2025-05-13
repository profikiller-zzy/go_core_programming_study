package main

func uniquePaths1(m int, n int) int {
	// 初始化dp，dp[i][j]表示从起点到达(i,j)的路径数
	var dp = make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// 初始化第一行和第一列
	for row := 0; row < m; row++ {
		dp[row][0] = 1
	}
	for col := 0; col < n; col++ {
		dp[0][col] = 1
	}
	// 动态规划计算路径数
	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			dp[row][col] = dp[row-1][col] + dp[row][col-1]
		}
	}
	return dp[m-1][n-1]
}

func uniquePaths(m int, n int) int {
	// 初始化dp，dp[i][j]表示从起点到达(i,j)的路径数
	var dp = make([]int, n)

	// 初始化第一行和第一列
	for col := 0; col < n; col++ {
		dp[col] = 1
	}
	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			dp[col] += dp[col-1]
		}
	}
	return dp[n-1]
}
