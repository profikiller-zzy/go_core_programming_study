package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.cn/problems/valid-sudoku/
func isValidSudoku(board [][]byte) bool {
	boxHash, rowHash, columnHash := make([][]bool, 9), make([][]bool, 9), make([][]bool, 9)
	for i := 0; i < 9; i++ {
		boxHash[i] = make([]bool, 9)
		rowHash[i] = make([]bool, 9)
		columnHash[i] = make([]bool, 9)
	}
	for row := 0; row < len(board); row++ {
		for column := 0; column < len(board[0]); column++ {
			if board[row][column] == '.' {
				continue
			} else {
				num, _ := strconv.Atoi(string(board[row][column]))
				if boxHash[row/3*3+column/3][num-1] || rowHash[row][num-1] || columnHash[column][num-1] {
					return false
				} else {
					boxHash[row/3*3+column/3][num-1] = true
					rowHash[row][num-1] = true
					columnHash[column][num-1] = true
				}
			}
		}
	}
	return true
}

func main() {
	sudoku := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(sudoku))
}
