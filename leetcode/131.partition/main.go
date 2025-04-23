package main

import "fmt"

func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}
	
	var result [][]string
	// 动态规划数组，dp[i][j]表示s[i:j+1]是否是回文串
	var dp = make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := 0; i < len(s); i++ {
		dp[i][i] = true // 单个字符是回文串
	}
	// 动态规划预处理，判断是否是回文串
	for length := 1; length <= len(s); length++ {
		for start := 0; start+length-1 < len(s); start++ {
			end := start + length - 1
			if s[start] == s[end] && (length <= 2 || dp[start+1][end-1]) {
				dp[start][end] = true
			}
		}
	}

	isPalindrome := func(start, end int) bool {
		// 判断s[start:end+1]是否是回文串
		if start > end {
			return false
		} else {
			return dp[start][end]
		}
	}

	var dfs func(int, []string)
	dfs = func(start int, path []string) {
		if start == len(s) {
			copyPath := make([]string, len(path))
			copy(copyPath, path)
			result = append(result, copyPath) // 递归结束，添加到结果中
			return
		}
		// start表示当前的起始位置， path表示已经分割的回文串
		for i := start; i < len(s); i++ {
			if isPalindrome(start, i) {
				// 递归调用
				path = append(path, s[start:i+1])
				dfs(i+1, path)
				path = path[:len(path)-1] // 回溯
			}
		}
	}
	dfs(0, []string{})
	return result
}

func main() {
	// 测试代码
	s := "abbab"
	result := partition(s)
	fmt.Println(result)
}
