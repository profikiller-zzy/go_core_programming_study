package main

import (
	"fmt"
)

func largestRectangleArea1(heights []int) int {
	n := len(heights)
	leftStack, rightStack := make([]int, 1), make([]int, 1)
	leftStack[0], rightStack[0] = -1, n
	left, right := make([]int, n), make([]int, n)

	for index := 0; index < n; index++ {
		// 维护一个单调递增栈
		// 如果当前柱子 height[index] <= 栈顶柱子的高度，说明栈顶柱子不可能是 index 左边第一个比它矮的
		// 而且对于后面的柱子来说，栈顶柱子也被 index 挡住了（因为 index 更矮或相等且更靠右），所以弹出栈顶
		for len(leftStack) > 1 && heights[index] <= heights[leftStack[len(leftStack)-1]] {
			leftStack = leftStack[:len(leftStack)-1]
		}
		left[index] = leftStack[len(leftStack)-1]
		leftStack = append(leftStack, index)
	}
	for index := n - 1; index >= 0; index-- {
		for len(rightStack) > 1 && heights[index] <= heights[rightStack[len(rightStack)-1]] {
			rightStack = rightStack[:len(rightStack)-1]
		}
		right[index] = rightStack[len(rightStack)-1]
		rightStack = append(rightStack, index)
	}

	var maxArea int
	for index := 0; index < n; index++ {
		curArea := (right[index] - left[index] - 1) * heights[index]
		if curArea > maxArea {
			maxArea = curArea
		}
	}
	return maxArea
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	leftStack, rightStack := make([]int, 1), make([]int, 1)
	leftStack[0], rightStack[0] = -1, n
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		// 维护一个单调递增栈
		// 如果当前柱子 height[index] <= 栈顶柱子的高度，说明栈顶柱子不可能是 index 左边第一个比它矮的
		// 而且对于后面的柱子来说，栈顶柱子也被 index 挡住了（因为 index 更矮或相等且更靠右），所以弹出栈顶
		for len(leftStack) > 1 && heights[i] <= heights[leftStack[len(leftStack)-1]] {
			leftStack = leftStack[:len(leftStack)-1]
		}
		left[i] = leftStack[len(leftStack)-1]
		leftStack = append(leftStack, i)
	}
	for i := n - 1; i >= 0; i-- {
		for len(rightStack) > 1 && heights[i] <= heights[rightStack[len(rightStack)-1]] {
			rightStack = rightStack[:len(rightStack)-1]
		}
		right[i] = rightStack[len(rightStack)-1]
		rightStack = append(rightStack, i)
	}

	var curArea, maxArea int
	for i := 0; i < n; i++ {
		curArea = (right[i] - left[i] - 1) * heights[i]
		if curArea > maxArea {
			maxArea = curArea
		}
	}
	return maxArea
}

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(heights))
}
