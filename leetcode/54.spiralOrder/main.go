package main

// https://leetcode.cn/problems/spiral-matrix/description

func spiralOrder(matrix [][]int) []int {
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

func main() {

}
