package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates) // 关键点1：排序用于剪枝和去重
	var path []int
	var result [][]int

	var dfs func(start int, sum int)
	dfs = func(start int, sum int) {
		if sum == target {
			// 创建副本保存结果
			newPath := make([]int, len(path))
			copy(newPath, path)
			result = append(result, newPath)
			return
		}

		for i := start; i < len(candidates); i++ { // 控制选择起点
			if sum+candidates[i] > target {
				break // 利用排序后的特性提前剪枝
			}

			// 递归时保持选择范围在[i, n)
			path = append(path, candidates[i])
			dfs(i, sum+candidates[i])
			path = path[:len(path)-1]
		}
	}

	dfs(0, 0)
	return result
}

func main() {}
