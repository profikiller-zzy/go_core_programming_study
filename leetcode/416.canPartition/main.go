package main

import (
	"fmt"
)

func canPartition(nums []int) bool {
	var (
		sum    int // 数组和
		maxNum int // 最大数
		n      int // 数组长度
	)
	n = len(nums)
	for _, num := range nums {
		sum += num
		if num > maxNum {
			maxNum = num
		}
	}
	if sum%2 != 0 || maxNum > sum/2 {
		return false
	}

	target := sum / 2
	// 其中 dp[i][j] 表示从数组的 [0,i] 下标范围内选取若干个正整数（可以是 0 个），是否存在一种选取方案使得被选取的正整数的和等于 j
	var dp = make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, target+1)
	}
	// 初始化
	for i := 0; i < n; i++ {
		dp[i][0] = true
	}
	dp[0][nums[0]] = true
	for row := 0; row < n; row++ {
		for col := 0; col <= target; col++ {
			if dp[row][col] || row == 0 {
				continue
			}
			if nums[row] > col { // 当前数字大于目标值，则不要选取当前值
				dp[row][col] = dp[row-1][col]
			} else { // 当前数字大于等于目标值，则选取当前值
				dp[row][col] = dp[row-1][col] || dp[row-1][col-nums[row]]
			}
		}
	}
	return dp[n-1][target]
}

func canPartition1(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2

	dp := make([]bool, target+1)
	dp[0] = true

	for _, num := range nums {
		// 注意：一维数组需要从后向前更新，防止重复使用同一个元素
		// dp[j] 表示从数组中是否能够选取一个子集，它的和恰好等于 j
		// 假设处理了之前的数得到 dp[j-num] 为 true，处理当前 num 时，说明可以选取一个子集(从当前num以及之前的数中选)，它的和恰好等于 j
		for j := target; j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}

	return dp[target]
}

func main() {
	nums := []int{3, 3, 3, 4, 5}
	fmt.Println(canPartition1(nums)) // 输出 true
}
