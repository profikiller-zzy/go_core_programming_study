package main

func solve(board [][]byte) {
	rows, cols := len(board), len(board[0])
	type pair struct {
		row int
		col int
	}

	var dfs func(int, int)
	dfs = func(x, y int) {
		board[x][y] = 'E' // 标记为边界连通的区域，临时修改
		if x > 0 && board[x-1][y] == 'O' {
			dfs(x-1, y)
		}
		if x <= rows-2 && board[x+1][y] == 'O' {
			dfs(x+1, y)
		}
		if y > 0 && board[x][y-1] == 'O' {
			dfs(x, y-1)
		}
		if y <= cols-2 && board[x][y+1] == 'O' {
			dfs(x, y+1)
		}
	}
	for row := 0; row < rows; row++ {
		if board[row][0] == 'O' {
			dfs(row, 0)
		}
		if board[row][cols-1] == 'O' {
			dfs(row, cols-1)
		}
	}
	for col := 0; col < cols; col++ {
		if board[0][col] == 'O' {
			dfs(0, col)
		}
		if board[rows-1][col] == 'O' {
			dfs(rows-1, col)
		}
	}
	// 遍历整个棋盘，翻转被包围的区域，恢复边界连通的区域
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if board[row][col] == 'O' {
				board[row][col] = 'X' // 被包围的 O -> X
			} else if board[row][col] == 'E' {
				board[row][col] = 'O' // 边界连通的 E -> O
			}
		}
	}
}
