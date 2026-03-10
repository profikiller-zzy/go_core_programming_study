package main

import "fmt"

func solveNQueens(n int) [][]string {
	var result [][]string
	path := make([][]byte, n)
	for i := 0; i < n; i++ {
		path[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			path[i][j] = '.'
		}
	}

	colLine, _45Line, _135Line := make([]bool, n), make([]bool, 2*n-1), make([]bool, 2*n-1)
	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			temp := make([]string, n)
			for i := 0; i < n; i++ {
				temp[i] = string(path[i])
			}
			result = append(result, temp)
		}
		for col := 0; col < n; col++ {
			if !colLine[col] && !_45Line[row+col] && !_135Line[col-row+n-1] {
				path[row][col] = 'Q'
				colLine[col] = true
				_45Line[row+col] = true
				_135Line[col-row+n-1] = true
				dfs(row + 1)
				path[row][col] = '.'
				colLine[col] = false
				_45Line[row+col] = false
				_135Line[col-row+n-1] = false
			}
		}
	}
	dfs(0)
	return result
}

func main() {
	result := solveNQueens(4)
	fmt.Println(result)
}
