package main

func maximalSquare1(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	var maxSlide int
	if m == 0 || n == 0 {
		return maxSlide
	}
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if matrix[row][col] == '1' {
				maxSlide = max(maxSlide, 1)
				curMaxSlide := min(m-row, n-col)
				for step := 1; step < curMaxSlide; step++ {
					flag := true
					if matrix[row+step][col+step] == '0' {
						break
					}
					for i := 0; i < step; i++ {
						if matrix[row+step][col+i] == '0' || matrix[row+i][col+step] == '0' {
							flag = false
							break
						}
					}
					if flag {
						maxSlide = max(maxSlide, step+1)
					} else {
						break
					}
				}
			}
		}
	}
	return maxSlide * maxSlide
}

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	if m == 0 || n == 0 {
		return 0
	}
	// dp[i][j]表示的是以matrix[i][j]为右下角并且内容全1的正方形的最大周长
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	var maxSlide int
	for row := 0; row < m; row++ {
		if matrix[row][0] == '1' {
			dp[row][0] = 1
			maxSlide = 1
		}
	}
	for col := 0; col < n; col++ {
		if matrix[0][col] == '1' {
			dp[0][col] = 1
			maxSlide = 1
		}
	}
	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			if matrix[row][col] == '0' {
				dp[row][col] = 0
			} else {
				dp[row][col] = min(dp[row-1][col], dp[row-1][col-1], dp[row][col-1]) + 1
				maxSlide = max(dp[row][col], maxSlide)
			}
		}
	}
	return maxSlide * maxSlide
}
