package main

type pair struct {
	x, y int
}

func totalNQueens(n int) int {
	if n == 1 {
		return 1
	}
	// 已经放入的皇后
	queens := make([]pair, 0)
	var total int
	var dfs func(int)
	dfs = func(row int) {
		if row == n {
			total++
			return
		}
		for col := 0; col < n; col++ {
			if isValid(row, col, queens) { // 这个位置合法，放入一个皇后
				queens = append(queens, pair{row, col})
				dfs(row + 1)
				queens = queens[:len(queens)-1] // 回溯
			}
		}
		return // 减枝
	}
	dfs(0)
	return total
}

func isValid(row, col int, queens []pair) bool {
	for index := 0; index < len(queens); index++ {
		if queens[index].y == col ||
			queens[index].x+queens[index].y == row+col ||
			row-queens[index].x == col-queens[index].y {
			return false
		}
	}
	return true
}
