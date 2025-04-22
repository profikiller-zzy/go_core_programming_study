package main

func subsets(nums []int) [][]int {
	var result [][]int
	result = append(result, []int{})
	if len(nums) == 0 {
		return result
	}

	var dfs func([]int, int)
	dfs = func(curRes []int, index int) {
		for i := index; i < len(nums); i++ {
			// 选择当前元素
			curRes = append(curRes, nums[i])
			result = append(result, append([]int{}, curRes...))
			// 递归选择下一个元素
			dfs(curRes, i+1)
			// 回溯，去掉当前元素
			curRes = curRes[:len(curRes)-1]
		}
	}
	dfs([]int{}, 0)
	return result
}

func main() {}
