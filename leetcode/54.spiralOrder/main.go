package main

// https://leetcode.cn/problems/spiral-matrix/description

func spiralOrder1(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	rowCount, colCount := len(matrix), len(matrix[0])
	visited := make([][]bool, rowCount)
	for i := range visited {
		visited[i] = make([]bool, colCount)
	}

	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 右、下、左、上
	dirIndex := 0
	row, col := 0, 0
	result := make([]int, 0, rowCount*colCount)

	for i := 0; i < rowCount*colCount; i++ {
		result = append(result, matrix[row][col])
		visited[row][col] = true

		nextRow := row + directions[dirIndex][0]
		nextCol := col + directions[dirIndex][1]

		if nextRow >= 0 && nextRow < rowCount && nextCol >= 0 && nextCol < colCount && !visited[nextRow][nextCol] {
			row = nextRow
			col = nextCol
		} else {
			// 改变方向
			dirIndex = (dirIndex + 1) % 4
			row += directions[dirIndex][0]
			col += directions[dirIndex][1]
		}
	}

	return result
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	var (
		rowCount, colCount       = len(matrix), len(matrix[0])
		result                   = make([]int, rowCount*colCount)
		index                    = 0
		left, right, top, bottom = 0, colCount - 1, 0, rowCount - 1
	)

	for left <= right && top <= bottom {
		for col := left; col <= right; col++ {
			result[index] = matrix[top][col]
			index++
		}
		for row := top + 1; row <= bottom; row++ {
			result[index] = matrix[row][right]
			index++
		}
		if right > left && bottom > top { // 防止重复扫描底部和左边的部分
			for col := right - 1; col >= left; col-- {
				result[index] = matrix[bottom][col]
				index++
			}
			for row := bottom - 1; row >= top+1; row-- {
				result[index] = matrix[row][left]
				index++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return result
}

func main() {

}
