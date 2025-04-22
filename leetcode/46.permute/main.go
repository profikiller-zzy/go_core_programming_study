package main

func permute(nums []int) [][]int {
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
