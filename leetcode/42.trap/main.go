package main

import "fmt"

// 下面这种方法本质上是通过递归查找左右边界，然后计算中间的面积差来得到可以装水的面积。这种写法是 暴力递归 + 局部扫描，导致性能较差，最坏情况下时间复杂度接近 O(n²)，所以容易 超时。

func fewer(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func returnBothEnds(start, end int, height []int) (int, int) {
	left, right := start, end
	for left < right-1 {
		isModified := false
		for i := left; i < right; i++ {
			if height[i] > height[left] {
				left = i
				isModified = true
				break
			}
		}
		for i := right; i > left; i-- {
			if height[i] > height[right] {
				right = i
				isModified = true
				break
			}
		}
		if !isModified {
			break
		}
	}
	return left, right
}

func trap1(height []int) int {
	totalArea := 0
	calculateArea := func(start, end int) {
		if end-start <= 1 {
			return
		} else {
			totalArea += fewer(height[start], height[end]) * (end - start - 1)
			for i := start + 1; i < end; i++ {
				totalArea -= height[i]
			}
		}
	}

	var returnArea func(start, end int)
	returnArea = func(start, end int) {
		if end-start <= 1 {
			return
		}
		left, right := returnBothEnds(start, end, height)
		calculateArea(left, right)
		returnArea(start, left)
		returnArea(right, end)
	}
	returnArea(0, len(height)-1)
	return totalArea
}

// trap 我们用两个指针 left 和 right 分别从左右向中间走。维护两个变量 leftMax 和 rightMax 表示左右的最大高度。我们每次移动较低的一侧，计算当前位置可以接多少水。
func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	left, right := 0, n-1
	leftMax, rightMax := 0, 0
	water := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				water += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				water += rightMax - height[right]
			}
			right--
		}
	}

	return water
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
}
