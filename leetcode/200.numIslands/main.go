package main

func numIslands(grid [][]byte) int {
	var count int // 岛屿数量
	rows, cols := len(grid), len(grid[0])
	var dfs func(int, int)
	dfs = func(x, y int) {
		// 从(x, y)开始进行深度优先搜索，将所有相连的1都标记为0
		grid[x][y] = '0'
		if x-1 >= 0 && grid[x-1][y] == '1' {
			dfs(x-1, y)
		}
		if x+1 < rows && grid[x+1][y] == '1' {
			dfs(x+1, y)
		}
		if y-1 >= 0 && grid[x][y-1] == '1' {
			dfs(x, y-1)
		}
		if y+1 < cols && grid[x][y+1] == '1' {
			dfs(x, y+1)
		}
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '1' {
				// 进行深度优先搜索
				count++
				dfs(row, col)
			}
		}
	}
	return count
}

func main() {

}
