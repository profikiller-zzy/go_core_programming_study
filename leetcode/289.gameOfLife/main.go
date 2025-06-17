package main

func gameOfLife(board [][]int) {
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
