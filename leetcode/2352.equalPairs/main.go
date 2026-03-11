package main

func equalPairs(grid [][]int) int {
	n := len(grid)
	compareRowAndCol := func(row, col int) bool {
		for i := 0; i < n; i++ {
			if grid[row][i] != grid[i][col] {
				return false
			}
		}
		return true
	}

	var count int
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if compareRowAndCol(row, col) {
				count++
			}
		}
	}
	return count
}
