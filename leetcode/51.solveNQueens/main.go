package main

import "fmt"

func solveNQueens(n int) [][]string {
	var result [][]string

	var curAns = make([][]byte, n) // 存储当前回溯的解
	for i := 0; i < n; i++ {
		curAns[i] = make([]byte, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			curAns[i][j] = '.'
		}
	}

	colPos := make([]bool, n)    // 记录每一列是否有皇后
	diag1 := make([]bool, 2*n-1) // 45°线上是否有皇后
	diag2 := make([]bool, 2*n-1) // 135°线上是否有皇后

	var isValid func(int, int) bool
	isValid = func(row, col int) bool {
		// 判断当前的皇后是否会被攻击
		if colPos[col] || diag1[row+col] || diag2[row-col+n-1] {
			return false
		}
		return true
	}

	var dfs func(int, [][]byte)
	dfs = func(row int, curAns [][]byte) {
		// 皇后是一行行从上往下落位的
		if row == n {
			// 找到了一种解法
			var ans []string
			for _, rowBytes := range curAns {
				ans = append(ans, string(rowBytes))
			}
			result = append(result, ans)
			return
		}
		for col := 0; col < n; col++ {
			// 判断这个位置是否会被已经落位的皇后攻击
			if isValid(row, col) {
				// 进行落位
				curAns[row][col] = 'Q'
				colPos[col] = true
				diag1[row+col] = true
				diag2[row-col+n-1] = true
				dfs(row+1, curAns)
				// 回溯
				curAns[row][col] = '.'
				colPos[col] = false
				diag1[row+col] = false
				diag2[row-col+n-1] = false
			}
		}
		return
	}
	dfs(0, curAns)
	return result
}

func main() {
	result := solveNQueens(4)
	fmt.Println(result)
}
