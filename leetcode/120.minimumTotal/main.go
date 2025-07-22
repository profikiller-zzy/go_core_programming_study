package main

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	dp := make([][]int, len(triangle))
	for index := 0; index < len(triangle); index++ {
		dp[index] = make([]int, index+1)
	}
	dp[0][0] = triangle[0][0]
	for row := 1; row < len(triangle); row++ {
		dp[row][0] = dp[row-1][0] + triangle[row][0]
	}

	for row := 1; row < len(triangle); row++ {
		for col := 1; col <= row; col++ {
			if col == row {
				dp[row][col] = dp[row-1][col-1] + triangle[row][col]
				continue
			}
			dp[row][col] = min(dp[row-1][col], dp[row-1][col-1]) + triangle[row][col]
		}
	}
	minSum := dp[len(triangle)-1][0]
	for col := 1; col < len(triangle); col++ {
		if dp[len(triangle)-1][col] < minSum {
			minSum = dp[len(triangle)-1][col]
		}
	}
	return minSum
}

func main() {
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	result := minimumTotal(triangle)
	println(result) // Output: 11
}
