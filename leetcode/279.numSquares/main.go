package main

func numSquares(n int) int {
	var dp = make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		// 求 dp[i]
		var minSteps = i
		for j := 1; j*j <= i; j++ {
			minSteps = min(minSteps, dp[i-j*j])
		}
		dp[i] = minSteps + 1
	}
	return dp[n]
}

func main() {
	n := 12
	println(numSquares(n))
}
