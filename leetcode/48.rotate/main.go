package main

import "fmt"

// https://leetcode.cn/problems/rotate-image/description

func rotate(matrix [][]int) {
	var n, depth int
	n = len(matrix)
	if n <= 1 {
		return
	}
	depth = (n + 1) / 2 // 需要处理多少圈

	for curDepth := 0; curDepth < depth; curDepth++ {
		for index := curDepth; index < n-1-curDepth; index++ {
			// pos0 matrix[curDepth][index] pos1 matrix[index][n-1-curDepth] pos2 matrix[n-1-curDepth][n-1-index] pos3 matrix[n-1-index][curDepth]
			// 顺时针旋转
			pos0, pos1, pos2, pos3 := matrix[curDepth][index], matrix[index][n-1-curDepth], matrix[n-1-curDepth][n-1-index], matrix[n-1-index][curDepth]
			matrix[curDepth][index], matrix[index][n-1-curDepth], matrix[n-1-curDepth][n-1-index], matrix[n-1-index][curDepth] = pos3, pos0, pos1, pos2
		}
	}
	return
}

func main() {
	matrix := make([][]int, 3)
	matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
	fmt.Println(matrix)
}
