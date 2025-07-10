package main

import "fmt"

func permute1(nums []int) [][]int {
	isVisited := make([]bool, len(nums))
	var result [][]int
	var permuteHelper func([]int)
	permuteHelper = func(curRes []int) {
		// 终止条件：当前排列长度等于原数组长度
		if len(curRes) == len(nums) {
			tmp := make([]int, len(curRes))
			copy(tmp, curRes)
			result = append(result, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if !isVisited[i] {
				isVisited[i] = true
				// 传递新切片，不修改当前层的 curRes
				permuteHelper(append(curRes, nums[i]))
				isVisited[i] = false // 回溯
			}
		}
	}

	permuteHelper([]int{})
	return result
}

func permute(nums []int) [][]int {
	var res [][]int
	if len(nums) == 0 {
		return res
	}

	var path []int
	isVisited := make([]bool, len(nums))
	var dfs func(int)
	dfs = func(index int) {
		if index == len(nums) {
			newPath := make([]int, len(path))
			copy(newPath, path)
			res = append(res, newPath)
			return
		}
		for i := 0; i < len(nums); i++ {
			if !isVisited[i] {
				path = append(path, nums[i])
				isVisited[i] = true
				dfs(index + 1)
				isVisited[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return res
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}
