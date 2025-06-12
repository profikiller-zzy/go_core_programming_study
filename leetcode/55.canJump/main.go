package main

import "fmt"

func canJump1(nums []int) bool {
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

// canJump 贪心算法解题
func canJump(nums []int) bool {
	var farestDis int
	for index := 0; index < len(nums); index++ {
		if index <= farestDis {
			farestDis = max(farestDis, nums[index]+index)
			if farestDis >= len(nums)-1 {
				return true
			}
		}
	}
	return false
}

func main() {
	nums := []int{2, 0, 0}
	fmt.Println(canJump(nums))
}
