package main

func minPathSum1(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	var dp = make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
	}
	// 初始化第一行和第一列
	dp[0][0] = grid[0][0]
	for row := 1; row < rows; row++ {
		dp[row][0] = grid[row][0] + dp[row-1][0]
	}
	for col := 1; col < cols; col++ {
		dp[0][col] = grid[0][col] + dp[0][col-1]
	}
	for row := 1; row < rows; row++ {
		for col := 1; col < cols; col++ {
			dp[row][col] = grid[row][col] + min(dp[row-1][col], dp[row][col-1])
		}
	}
	return dp[rows-1][cols-1]
}
