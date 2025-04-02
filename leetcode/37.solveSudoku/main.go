package main

func solveSudoku(board [][]byte) {
	backtrack(board)
}

func backtrack(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for c := byte('1'); c <= byte('9'); c++ {
					if isValid(board, i, j, c) { // 填入c是否合法
						board[i][j] = c
						if backtrack(board) {
							return true
						}
						board[i][j] = '.' // 回溯
					}
				}
				return false // 当前空格无法填入任何数字，触发回溯
			}
		}
	}
	return true // 所有空格已填满
}

func isValid(board [][]byte, row, col int, c byte) bool {
	// 检查当前行是否存在相同数字
	for i := 0; i < 9; i++ {
		if board[row][i] == c {
			return false
		}
	}

	// 检查当前列是否存在相同数字
	for i := 0; i < 9; i++ {
		if board[i][col] == c {
			return false
		}
	}

	// 检查3x3的小方格是否存在相同数字
	startRow, startCol := (row/3)*3, (col/3)*3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == c {
				return false
			}
		}
	}

	return true
}

func main() {
	// 示例输入，此处省略，可根据实际需要添加测试用例
}
