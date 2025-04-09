package main

import "fmt"

// https://leetcode.cn/problems/longest-consecutive-sequence/description
func longestConsecutive(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	// 使用map来存储数字
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}
	maxLength := 0
	for _, num := range nums {
		if !numSet[num] {
			continue
		}
		numSet[num] = false
		curLength := 1
		// 从num前后开始查找连续的数字，并将其标记为false
		for next := num + 1; numSet[next]; next++ {
			curLength++
			numSet[next] = false
		}
		for prev := num - 1; numSet[prev]; prev-- {
			curLength++
			numSet[prev] = false
		}
		if curLength > maxLength {
			maxLength = curLength
		}
	}
	return maxLength
}

func main() {
	nums := []int{1, 0, 1, 2}
	fmt.Println(longestConsecutive(nums))
}
