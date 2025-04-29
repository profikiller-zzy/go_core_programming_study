package main

import "container/list"

type Stack struct {
	len  int
	data []int
}

func newStack() *Stack {
	return &Stack{
		len:  0,
		data: make([]int, 0),
	}
}

func (s *Stack) push(c int) {
	s.data = append(s.data, c)
	s.len++
}

func (s *Stack) pop() {
	if s.len == 0 {
		return
	}
	s.data = s.data[:s.len-1]
	s.len--
}

func (s *Stack) top() int {
	if s.len == 0 {
		return -1
	}
	return s.data[s.len-1]
}

func (s *Stack) isEmpty() bool {
	return s.len == 0
}

// largestRectangleArea1 参考leetcode官方题解单调栈解决最大矩形面积问题
// 使用自定义栈实现
func largestRectangleArea1(heights []int) int {
	n := len(heights)
	left, right := newStack(), newStack()
	left.push(-1)
	right.push(n)
	leftIndex, rightIndex := make([]int, n), make([]int, n)

	for index := 0; index < n; index++ {
		for left.len > 1 && heights[index] <= heights[left.top()] {
			left.pop()
		}
		leftIndex[index] = left.top()
		left.push(index)
	}
	for index := n - 1; index >= 0; index-- {
		for right.len > 1 && heights[index] <= heights[right.top()] {
			right.pop()
		}
		rightIndex[index] = right.top()
		right.push(index)
	}
	maxArea := 0
	for index := 0; index < n; index++ {
		curArea := (rightIndex[index] - leftIndex[index] - 1) * heights[index]
		if curArea > maxArea {
			maxArea = curArea
		}
	}
	return maxArea
}

// largestRectangleArea 参考leetcode官方题解单调栈解决最大矩形面积问题
// 使用container/list实现
func largestRectangleArea(heights []int) int {
	n := len(heights)
	left, right := list.New(), list.New()
	left.PushBack(-1)
	right.PushBack(n)
	leftIndex, rightIndex := make([]int, n), make([]int, n)

	for index := 0; index < n; index++ {
		for left.Len() > 1 && heights[index] <= heights[left.Back().Value.(int)] {
			left.Remove(left.Back())
		}
		leftIndex[index] = left.Back().Value.(int)
		left.PushBack(index)
	}
	for index := n - 1; index >= 0; index-- {
		for right.Len() > 1 && heights[index] <= heights[right.Back().Value.(int)] {
			right.Remove(right.Back())
		}
		rightIndex[index] = right.Back().Value.(int)
		right.PushBack(index)
	}
	maxArea := 0
	for index := 0; index < n; index++ {
		curArea := (rightIndex[index] - leftIndex[index] - 1) * heights[index]
		if curArea > maxArea {
			maxArea = curArea
		}
	}
	return maxArea
}
