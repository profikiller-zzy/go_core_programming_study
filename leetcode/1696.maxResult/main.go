package main

import "fmt"

func maxResult(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]   // dp 动态规划，表示到达每个位置的最大得分，dp[j] = max(dp[i]) + nums[j]，其中 i 在 [j-k, j-1] 范围内
	queue := []int{0} // 使用双端队列来维护当前滑动窗口中的最大值的索引

	push := func(index int) { // push 操作，添加索引到队列右侧，不断将小于当前值的索引出队，直到队列为空或者队列尾的值大于等于当前值
		for len(queue) > 0 && dp[queue[len(queue)-1]] < dp[index] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, index)
	}

	for index := 1; index < n; index++ {
		for queue[0] < index-k { // 移除不在滑动窗口内的索引
			queue = queue[1:]
		}
		dp[index] = nums[index] + dp[queue[0]]
		push(index)
	}
	return dp[n-1] // 返回到达最后一个位置的最大得分
}

func main() {
	nums := []int{10, -5, -2, 4, 0, 3}
	fmt.Println(maxResult(nums, 3))
}
