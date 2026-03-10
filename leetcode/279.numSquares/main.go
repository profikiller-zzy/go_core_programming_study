package main

func numSquares1(n int) int {
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

// numSquares 广度优先搜索解题
func numSquares(n int) int {
	queue := []int{n}
	visited := make([]bool, n+1)
	visited[n] = true
	step := 0
	for len(queue) > 0 {
		size := len(queue)
		step++
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			for j := 1; j*j <= cur; j++ {
				next := cur - j*j
				if next == 0 {
					return step
				}
				if !visited[next] {
					queue = append(queue, next)
					visited[next] = true
				}
			}
		}
	}
	return n
}

func main() {
	n := 12
	println(numSquares(n))
}
