package main

import "fmt"

func findMaxAverage(nums []int, k int) float64 {
	var curSum, maxSum int
	for i := 0; i < k; i++ {
		curSum += nums[i]
	}
	maxSum = curSum
	for i := k; i < len(nums); i++ {
		curSum -= nums[i-k]
		curSum += nums[i]
		if maxSum < curSum {
			maxSum = curSum
		}
	}
	return float64(maxSum) / float64(k)
}

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	fmt.Println(findMaxAverage(nums, 4))
}
