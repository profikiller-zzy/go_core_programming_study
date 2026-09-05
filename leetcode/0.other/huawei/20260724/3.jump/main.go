package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

/*
题目描述：
有 n 块直线排列的地砖，小红从第一块地砖之前出发。每回合前进 1、2 或 3 块，并自动拾取落脚砖上的宝物。
她必须在不超过 m 回合内恰好到达最后一块地砖。求可携带的最大宝物价值，无法到达输出 -1。

输入描述：
第一行输入 n, m，第二行输入 n 个宝物价值。保证 5 <= n <= 10000，2 <= m <= 5000，价值在 [0, 5]。

输出描述：
输出最大价值，无法在回合限制内到达则输出 -1。

示例1
输入：5 3
     1 2 1 1 3
输出：6
说明：依次落在第 2、3 或 4、5 块地砖。

示例2
输入：10 3
     0 0 3 1 2 3 0 0 0 0
输出：-1
说明：三回合最多前进 9 格，无法从地窖外到达第 10 块。
*/

func solution(n, m int, path []int) int {
	if n > 3*m {
		return -1
	}
	maxSteps := m
	dp := make([][]int, maxSteps+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = -math.MaxInt32
		}
	}
	dp[0][0] = 0
	dp[1][1] = path[1]
	dp[1][2] = path[2]
	dp[1][3] = path[3]
	for step := 2; step <= maxSteps; step++ {
		left, right := step, 3*step     // 经过step回合后，最少能到step，最多能到3*step
		if left < n-(maxSteps-step)*3 { // step回合最近也得到n - (maxSteps-step)*3，不然后面跳不到位置n了
			left = n - (maxSteps-step)*3
		}
		if right > n {
			right = n
		}
		for pos := left; pos <= right; pos++ {
			best := -math.MaxInt32
			for i := 1; i <= 3; i++ {
				prevPos := pos - i
				if prevPos <= 0 {
					continue
				}
				if dp[step-1][prevPos] > 0 && dp[step-1][prevPos] > best {
					best = dp[step-1][prevPos]
				}
			}
			dp[step][pos] = best + path[pos]
		}
	}
	var result int
	for step := 1; step <= maxSteps; step++ {
		if dp[step][n] > result {
			result = dp[step][n]
		}
	}
	return result
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n, m int
	fmt.Fscan(in, &n, &m)
	path := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &path[i])
	}
	fmt.Fprintln(out, solution(n, m, path))
}
