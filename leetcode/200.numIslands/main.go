package main

func numIslands(grid [][]byte) int {
	var count int // 岛屿数量
	var dfs func(int, int)
	dfs = func(x, y int) {
		// 从(x, y)开始进行深度优先搜索，将所有相连的1都标记为0

	}

	for row := 0; row < len(grid)-1; row++ {
		for col := 0; col < len(grid[0])-1; col++ {
			if grid[row][col] == '1' {
				// 进行深度优先搜索
				dfs(row, col)
			}
		}
	}
}

func main() {

}
