package main

func searchMatrix(matrix [][]int, target int) bool {
	// 从右上角开始寻找
	rows, cols := len(matrix), len(matrix[0])
	row, col := 0, cols-1
	for row < rows && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if target < matrix[row][col] {
			col-- // 往左走
		} else {
			row++ // 往下走
		}
	}
	return false
}
