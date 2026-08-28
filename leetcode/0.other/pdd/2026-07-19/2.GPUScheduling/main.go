package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func minBatches(n, m, k, c int, data []int) int {
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})

	// next[i]表示从第i个请求开始，一个批次能处理到的位置，即[i, next[i])，左闭右开
	next := make([]int, n)
	right := 0
	for left := 0; left < n; left++ {
		if right < left {
			right = left
		}
		for right < n && data[right]-data[left] <= k && right-left < c {
			right++
		}
		next[left] = right
	}
	// dp[i][j] 为处理完排序后的前i个请求，恰好丢弃j个请求的时候，需要的最小批次数量
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = n + 1
		}
	}
	dp[0][0] = 0
	for i := 0; i < n; i++ {
		for j := 0; j <= m; j++ {
			if dp[i][j] == n+1 { // 该状态不可达，例如i < j的请求，最开始的可达状态只有dp[0][0]
				continue
			}
			// 选择1: 丢弃这个请求
			if j < m {
				dp[i+1][j+1] = min(dp[i+1][j+1], dp[i][j])
			}
			// 选择2: 从当前请求i开始选择一个批次
			dp[next[i]][j] = min(dp[next[i]][j], dp[i][j]+1)
		}
	}
	result := n + 1
	for j := 0; j <= m; j++ {
		result = min(result, dp[n][j])
	}
	return result
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)

	for testCase := 0; testCase < testCount; testCase++ {
		var n, m, k, c int
		fmt.Fscan(in, &n, &m, &k, &c)
		data := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &data[j])
		}
		fmt.Fprintln(out, minBatches(n, m, k, c, data))
	}
}
