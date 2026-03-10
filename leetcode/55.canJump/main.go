package main

import "fmt"

// canJump 贪心算法解题
func canJump(nums []int) bool {
	n := len(nums)
	maxReach := 0
	for i := 0; i <= maxReach && i < n; i++ {
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
		}
		if maxReach >= n-1 {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{2, 0, 0}
	fmt.Println(canJump(nums))
}
