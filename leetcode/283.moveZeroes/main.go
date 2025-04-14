package main

import "fmt"

// https://leetcode.cn/problems/move-zeroes/description
func moveZeroes1(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	leftMove := func(start, end int) {
		for i := start; i < end; i++ {
			nums[i] = nums[i+1]
		}
		nums[end] = 0
	}
	end := len(nums) - 1
	for i := 0; i < end; i++ {
		if nums[i] == 0 {
			leftMove(i, end)
			end--
			i--
		}
	}
}

// moveZeroes 双指针
func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		// 左指针指向当前已经处理好的序列的尾部，右指针指向待处理序列的头部。
		// 右指针不断向右移动，每次右指针指向非零数，则将左右指针对应的数交换，同时左指针右移。
		// 这样可以保证左指针左边的数都是非零数，右指针左边到左指针的数都是零。
		// 并且非零数的顺序不变。
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left] // 做一次交换
			left++
		}
		right++
	}
}

func main() {
	nums := []int{0, 0, 1}
	moveZeroes(nums)
	fmt.Println(nums)
}
