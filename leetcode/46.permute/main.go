package main

import "fmt"

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) == 0 {
		return res
	}

	path := make([]int, 0, len(nums))
	isVisited := make([]bool, len(nums))

	// 定义递归函数
	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			res = append(res, append([]int(nil), path...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if !isVisited[i] {
				path = append(path, nums[i])
				isVisited[i] = true
				dfs()
				isVisited[i] = false
				path = path[:len(path)-1] // 回溯
			}
		}
	}

	dfs()
	return res
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}
