package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	var path []int
	var result [][]int
	var dfs func(int, int, int)
	dfs = func(index int, start int, curSum int) {
		if index == k { // 现在已经挑选了k个数字
			if curSum == n {
				result = append(result, append([]int{}, path...))
			}
			return
		}
		if 9-start+1 < k-index {
			return // 剩下的数字不够选到k个数
		}
		for num := start; num <= 9; num++ {
			path = append(path, num)
			// 选择了num，那么下次挑选的数一定大于num，这是保证不重复
			dfs(index+1, num+1, curSum+num)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 1, 0)
	return result
}

func main() {
	fmt.Println(combinationSum3(9, 45))
}
