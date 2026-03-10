package main

import "fmt"

func partition(s string) [][]string {
	n := len(s)
	if n == 0 {
		return [][]string{}
	}

	// DP 预处理：dp[i][j] 表示 s[i:j+1] 是否是回文串
	// 使用一维切片数组，Go 中 [][]bool 实际上是多个独立的一维数组
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for j := 0; j < n; j++ { // 右边界增长
		for i := 0; i <= j; i++ {
			// 如果首尾字符相同，且中间部分也是回文（或长度<=2），则是回文
			dp[i][j] = s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1])
		}
	}

	var result [][]string
	var path []string
	var dfs func(start int)
	dfs = func(start int) {
		if start == n {
			// 必须复制 path，因为 path 是引用类型，后续回溯会修改它
			temp := make([]string, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		// 尝试所有可能的结束位置
		for end := start; end < n; end++ {
			if dp[start][end] {
				path = append(path, s[start:end+1])
				dfs(end + 1)
				path = path[:len(path)-1]
			}
		}
	}

	dfs(0)
	return result
}

func main() {
	// 测试代码
	s := "abbab"
	result := partition(s)
	fmt.Println(result)
}
