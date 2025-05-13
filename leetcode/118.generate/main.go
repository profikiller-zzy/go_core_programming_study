package main

func generate(numRows int) [][]int {
	var res = make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
		res[i][0] = 1
		res[i][i] = 1
		if i == 0 {
			continue
		}
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}
	return res
}
