package main

func countSquares(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	var count int
	for row := 0; row < m; row++ {
		if matrix[row][0] == 1 {
			dp[row][0] = 1
			count += 1
		}
	}
	for col := 1; col < n; col++ {
		if matrix[0][col] == 1 {
			dp[0][col] = 1
			count += 1
		}
	}
	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			if matrix[row][col] == 1 {
				dp[row][col] = min(dp[row-1][col], dp[row-1][col-1], dp[row][col-1]) + 1
				count += dp[row][col]
			}
		}
	}
	return count
}
