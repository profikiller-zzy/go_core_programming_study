package main

import "fmt"

// https://leetcode.cn/problems/search-a-2d-matrix-ii/description

func searchMatrix1(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	if rows == 0 || cols == 0 {
		return false
	}
	if matrix[0][0] == target {
		return true
	}
	if matrix[0][0] > target {
		return false
	}

	var curRow, curCol int
	var isLast bool
	for {
		if matrix[curRow][curCol] == target {
			return true
		}
		if matrix[curRow][curCol] > target {
			for row := 0; row < curRow; row++ {
				if curRow > curCol {
					break
				}
				if matrix[row][curCol] == target {
					return true
				}
			}
			for col := 0; col < curCol; col++ {
				if curCol > curRow {
					break
				}
				if matrix[curRow][col] == target {
					return true
				}
			}
		}
		if curRow < rows-1 {
			curRow++
		}
		if curCol < cols-1 {
			curCol++
		}
		if curRow == rows-1 && curCol == cols-1 {
			if !isLast {
				isLast = true
			} else {
				return false
			}
		}
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	// z 字形查找
	rows, cols := len(matrix), len(matrix[0])
	curRow, curCol := 0, cols-1
	for curRow < rows && curCol >= 0 {
		if matrix[curRow][curCol] == target {
			return true
		} else if matrix[curRow][curCol] > target {
			curCol--
		} else {
			curRow++
		}
	}
	return false
}

func main() {
	//matrix := [][]int{
	//	{1, 4, 7, 11, 15},
	//	{2, 5, 8, 12, 19},
	//	{3, 6, 9, 16, 22},
	//	{10, 13, 14, 17, 24},
	//	{18, 21, 23, 26, 30},
	//}
	matrix := [][]int{
		{2, 5},
		{2, 8},
		{7, 9},
		{7, 11},
		{9, 11},
	}
	target := 7
	fmt.Println(searchMatrix(matrix, target))
}
