package main

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
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
			if board[row][col] == word[0] {
				if dfs(row, col, 0) {
					return true
				}
			}
		}
	}
	return false
}
