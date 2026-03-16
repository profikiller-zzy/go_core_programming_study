package main

func nearestExit(maze [][]byte, entrance []int) int {
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	rows, cols := len(maze), len(maze[0])
	var step int
	queue := [][2]int{{entrance[0], entrance[1]}}
	maze[entrance[0]][entrance[1]] = '+'

	for len(queue) > 0 {
		step++
		size := len(queue)
		for i := 0; i < size; i++ {
			row, col := queue[0][0], queue[0][1]
			queue = queue[1:]
			for _, dir := range directions {
				nextRow, nextCol := row+dir[0], col+dir[1]
				if nextRow >= 0 && nextRow < rows && nextCol >= 0 && nextCol < cols && maze[nextRow][nextCol] == '.' {
					if nextRow == 0 || nextRow == rows-1 || nextCol == 0 || nextCol == cols-1 {
						return step // bfs 搜索第一次到达的出口就是最近的出口
					}
					queue = append(queue, [2]int{nextRow, nextCol})
					maze[nextRow][nextCol] = '+'
				}
			}
		}
	}

	return -1 // 没有找到，返回-1
}
