package main

import "fmt"

// https://leetcode.cn/problems/container-with-most-water/description

func maxArea(height []int) int {
	var fewer = func(a int, b int) int {
		if a > b {
			return b
		}
		return a
	}
	mArea := 0
	left, right := 0, len(height)-1
	for left < right {
		curArea := fewer(height[left], height[right]) * (right - left)
		if curArea > mArea {
			mArea = curArea
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return mArea
}

func main() {
	var height []int
	height = make([]int, 9)
	height = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
