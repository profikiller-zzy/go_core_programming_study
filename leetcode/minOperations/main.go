package main

import "fmt"

func main() {
	var grid [][]int = [][]int{
		{1, 5},
		{2, 3},
	}
	x := 1
	fmt.Println(minOperations(grid, x))
}

func minOperations(grid [][]int, x int) int {
	numRow := len(grid)
	numCol := len(grid[0])

	remain := make([]bool, x)
	// axis 用于记录每个点在数轴上的位置
	axis := make(map[int]int)
	var left, right int = (grid[0][0] % x) / x, (grid[0][0] % x) / x

	remain[grid[0][0]%x] = true
	for r := 0; r < numRow; r++ {
		for c := 0; c < numCol; c++ {
			// 余数不一样，怎么加减都不可能相等
			mod := grid[r][c] % x
			if !remain[mod] {
				return -1
			} else {
				grid[r][c] = (grid[r][c] - mod) / x
				axis[grid[r][c]] += 1
				if grid[r][c] < left {
					left = grid[r][c]
				}
				if grid[r][c] > right {
					right = grid[r][c]
				}
			}
		}
	}

	var totalSteps int
	for left != right {
		if axis[left] <= axis[right] {
			totalSteps += axis[left]
			axis[left+1] += axis[left]
			delete(axis, left)
			left++
		} else {
			totalSteps += axis[right]
			axis[right-1] += axis[right]
			delete(axis, right)
			right--
		}
	}
	return totalSteps
}
