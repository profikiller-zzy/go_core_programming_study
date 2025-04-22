package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates) // 关键点1：排序用于剪枝和去重
	var result [][]int

	var dfs func(start int, path []int, sum int)
	dfs = func(start int, path []int, sum int) {
		if sum == target {
			// 创建副本保存结果
			result = append(result, append([]int{}, path...))
			return
		}

		for i := start; i < len(candidates); i++ { // 控制选择起点
			if sum+candidates[i] > target {
				break // 利用排序后的特性提前剪枝
			}

			// 递归时保持选择范围在[i, n)
			dfs(i, append(path, candidates[i]), sum+candidates[i])
		}
	}

	dfs(0, []int{}, 0)
	return result
}

func main() {}
