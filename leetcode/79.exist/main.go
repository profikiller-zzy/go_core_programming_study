package main

func exist1(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	var dfs func(int, int, int) bool
	dfs = func(x, y, index int) bool {
		if index == len(word) {
			return true
		}
		if x < 0 || x >= rows || y < 0 || y >= cols || visited[x][y] || board[x][y] != word[index] {
			return false
		}

		visited[x][y] = true
		found := dfs(x-1, y, index+1) ||
			dfs(x+1, y, index+1) ||
			dfs(x, y-1, index+1) ||
			dfs(x, y+1, index+1)
		visited[x][y] = false // 回溯

		return found
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == word[0] && dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	var dfs func(int, int, int) bool
	dfs = func(row, col, index int) bool {
		if index == len(word) {
			return true
		}
		if row < 0 || row >= m || col < 0 || col >= n || board[row][col] != word[index] {
			return false
		}
		ch := board[row][col]
		board[row][col] = '#'
		dirs := [][2]int{
			{1, 0}, {-1, 0}, {0, 1}, {0, -1},
		}
		for _, dir := range dirs {
			nextRow, nextCol := row+dir[0], col+dir[1]
			if dfs(nextRow, nextCol, index+1) {
				return true
			}
		}
		board[row][col] = ch
		return false
	}
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if dfs(row, col, 0) {
				return true
			}
		}
	}
	return false
}
