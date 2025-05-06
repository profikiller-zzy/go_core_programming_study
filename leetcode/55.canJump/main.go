package main

import "fmt"

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	var dp = make([]bool, len(nums))
	dp[0] = true
	for index := 0; index < len(nums); index++ {
		if dp[index] {
			if index+nums[index] >= len(nums)-1 {
				return true
			}
			for i := index + nums[index]; i >= index+1; i-- {
				if dp[i] {
					break
				} else {
					dp[i] = true
				}
			}
		} else {
			break
		}
	}
	return false
}

func main() {
	nums := []int{2, 0, 0}
	fmt.Println(canJump(nums))
}
