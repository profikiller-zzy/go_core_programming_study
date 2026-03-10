package main

type Pos struct {
	x int
	y int
}

func orangesRotting(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	freshCount := 0
	var queue []Pos // BFS 队列

	// 初始化：统计新鲜橘子数量，记录腐烂橘子位置
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if grid[x][y] == 1 {
				freshCount++
			} else if grid[x][y] == 2 {
				queue = append(queue, Pos{x, y})
			}
		}
	}

	// 没有新鲜橘子，直接返回 0
	if freshCount == 0 {
		return 0
	}

	// 四个方向：上、下、左、右
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	steps := 0

	for len(queue) > 0 {
		levelSize := len(queue) // 当前层的腐烂橘子
		for i := 0; i < levelSize; i++ {
			pos := queue[0]
			queue = queue[1:]
			x, y := pos.x, pos.y

			// 检查四个方向
			for _, d := range directions {
				nx, ny := x+d[0], y+d[1]
				if nx >= 0 && nx < rows && ny >= 0 && ny < cols && grid[nx][ny] == 1 {
					// 腐烂新鲜橘子
					grid[nx][ny] = 2
					freshCount--
					queue = append(queue, Pos{nx, ny})
				}
			}
		}
		if len(queue) > 0 {
			steps++ // 每完成一层传播，步骤加 1
		}
	}

	// 检查是否还有新鲜橘子
	if freshCount > 0 {
		return -1
	}
	return steps
}

func main() {
}
