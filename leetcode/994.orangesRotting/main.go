package main

type Pos struct {
	x int
	y int
}

func orangesRotting(grid [][]int) int {
	// 1. 先遍历一遍，找到所有腐烂的橘子
	sideRotted := make([]Pos, 0)
	rows, cols := len(grid), len(grid[0])
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 2 {
				sideRotted = append(sideRotted, Pos{x: row, y: col})
			}
		}
	}
	var spread func() bool
	spread = func() bool {
		// 进行一次腐烂
		isSpread := false
		nextRotted := make([]Pos, 0)
		for _, pos := range sideRotted {
			x, y := pos.x, pos.y
			if x-1 >= 0 && grid[x-1][y] == 1 {
				isSpread = true
				grid[x-1][y] = 2
				nextRotted = append(nextRotted, Pos{x: x - 1, y: y})
			}
			if x+1 < rows && grid[x+1][y] == 1 {
				isSpread = true
				grid[x+1][y] = 2
				nextRotted = append(nextRotted, Pos{x: x + 1, y: y})
			}
			if y-1 >= 0 && grid[x][y-1] == 1 {
				isSpread = true
				grid[x][y-1] = 2
				nextRotted = append(nextRotted, Pos{x: x, y: y - 1})
			}
			if y+1 < cols && grid[x][y+1] == 1 {
				isSpread = true
				grid[x][y+1] = 2
				nextRotted = append(nextRotted, Pos{x: x, y: y + 1})
			}
		}
		sideRotted = nextRotted
		return isSpread
	}

	var step int
	for spread() {
		step++
	}
	// 检查是否还有新鲜橘子
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return step
}

func main() {
}
