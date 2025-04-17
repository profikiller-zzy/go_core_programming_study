package main

func numIslands(grid [][]byte) int {
	rows, cols := len(grid), len(grid[0])
	var count int
	var dfs func(int, int)
	dfs = func(row, col int) {
		// 深度优先搜索，将相邻的陆地标记为水
		grid[row][col] = byte('0')
		if row-1 >= 0 && grid[row-1][col] == '1' {
			dfs(row-1, col)
		}
		if row+1 < rows && grid[row+1][col] == '1' {
			dfs(row+1, col)
		}
		if col-1 >= 0 && grid[row][col-1] == '1' {
			dfs(row, col-1)
		}
		if col+1 < cols && grid[row][col+1] == '1' {
			dfs(row, col+1)
		}
	}
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '1' {
				count++
				dfs(row, col) // 调用深度优先搜索将相邻的陆地全部转化为水
			}
		}
	}
	return count
}
