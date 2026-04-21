package main

func gameOfLife1(board [][]int) {
	rows, cols := len(board), len(board[0])
	newBoard := make([][]int, rows)
	for index := 0; index < rows; index++ {
		newBoard[index] = make([]int, cols)
	}
	var livingCellNum func(int, int) int
	livingCellNum = func(row int, col int) int {
		var num int
		for r := row - 1; r <= row+1; r++ {
			for c := col - 1; c <= col+1; c++ {
				if r >= 0 && r < rows && c >= 0 && c < cols {
					num += board[r][c]
				}
			}
		}
		num -= board[row][col]
		return num
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			lcNum := livingCellNum(row, col)
			if board[row][col] == 0 {
				if lcNum == 3 {
					newBoard[row][col] = 1
				} else {
					newBoard[row][col] = 0
				}
			} else {
				if lcNum == 3 || lcNum == 2 {
					newBoard[row][col] = 1
				} else {
					newBoard[row][col] = 0
				}
			}
		}
	}
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			board[row][col] = newBoard[row][col]
		}
	}
}

// gameOfLife 思路：使用中间状态来节约存储空间
// -1: 由活细胞变成死细胞 2: 由死细胞复活为活细胞
func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	countLives := func(row, col int) int {
		var count int
		for x := row - 1; x <= row+1; x++ {
			for y := col - 1; y <= col+1; y++ {
				if (x == row && y == col) || x < 0 || x >= m || y < 0 || y >= n {
					continue
				}
				if board[x][y] == 1 || board[x][y] == -1 {
					count++
				}
			}
		}
		return count
	}
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			liveNeighbors := countLives(row, col)
			if board[row][col] == 1 && (liveNeighbors < 2 || liveNeighbors > 3) {
				board[row][col] = -1
			}
			if board[row][col] == 0 && liveNeighbors == 3 {
				board[row][col] = 2
			}
		}
	}
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if board[row][col] == -1 {
				board[row][col] = 0
			}
			if board[row][col] == 2 {
				board[row][col] = 1
			}
		}
	}
}
