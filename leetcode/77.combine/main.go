package main

import "fmt"

func combine(n int, k int) [][]int {
	var res [][]int
	if k == 0 {
		return res
	}

	var path []int
	var dfs func(int, int)
	dfs = func(index int, cur int) {
		if index == k {
			newPath := make([]int, len(path))
			copy(newPath, path)
			res = append(res, newPath) // 防止底层引用同一个地址
			return
		}
		for num := cur; num <= n; num++ {
			path = append(path, num)
			dfs(index+1, num+1)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 1)
	return res
}

func main() {
	fmt.Println(combine(4, 3))
}
