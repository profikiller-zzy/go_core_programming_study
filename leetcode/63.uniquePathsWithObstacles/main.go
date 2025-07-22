package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	rows, cols := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, rows)
	for row := 0; row < rows; row++ {
		dp[row] = make([]int, cols)
	}
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if obstacleGrid[row][col] == 1 {
				dp[row][col] = 0 // 遇到障碍物，路径数为0
			} else if row == 0 && col == 0 {
				dp[row][col] = 1 // 起点
			} else if row == 0 {
				dp[row][col] = dp[row][col-1] // 第一行只能从左边来
			} else if col == 0 {
				dp[row][col] = dp[row-1][col] // 第一列只能从上边来
			} else {
				dp[row][col] = dp[row-1][col] + dp[row][col-1] // 从上边和左边来的路径数之和
			}
		}
	}
	return dp[rows-1][cols-1]
}
