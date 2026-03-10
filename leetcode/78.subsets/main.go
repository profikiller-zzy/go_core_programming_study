package main

import "fmt"

func subsets(nums []int) [][]int {
	result := make([][]int, 0, 1<<len(nums)) // 预分配容量：2^N
	path := make([]int, 0, len(nums))        // 预分配 path 容量

	var dfs func(start int)
	dfs = func(start int) {
		result = append(result, append([]int{}, path...))

		for i := start; i < len(nums); i++ {
			// 选择
			path = append(path, nums[i])
			// 递归
			dfs(i + 1)
			// 回溯 (移除最后一个元素)
			path = path[:len(path)-1]
		}
	}

	dfs(0)
	return result
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
