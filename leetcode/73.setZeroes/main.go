package main

import "fmt"

// https://leetcode.cn/problems/set-matrix-zeroes/description

func setZeroes(matrix [][]int) {
	rows := make(map[int]bool)
	cols := make(map[int]bool)

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] == 0 {
				rows[row] = true
				cols[col] = true
			}
		}
	}

	for row, _ := range rows {
		for col := 0; col < len(matrix[0]); col++ {
			matrix[row][col] = 0
		}
	}

	for col, _ := range cols {
		for row := 0; row < len(matrix); row++ {
			matrix[row][col] = 0
		}
	}
}

func main() {
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix)
	fmt.Println(matrix)
}
