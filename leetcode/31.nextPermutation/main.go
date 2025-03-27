package main

import "fmt"

// https://leetcode.cn/problems/next-permutation/
func nextPermutation(nums []int) {
	reverse := func(start int, end int) {
		for i := start; i <= (start+end)/2; i++ {
			// 计算对称位置的索引
			j := end - (i - start)
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] { // 从后往前找到第一个升序的位置，这说明找到了需要调整的位置，即[i:len(nums)-1]
			for j := len(nums) - 1; j > i; j-- {
				if nums[j] > nums[i] { // 从后往前找到第一个比nums[i]大的位置，交换两个位置的值
					nums[i], nums[j] = nums[j], nums[i]
					reverse(i+1, len(nums)-1)
					return
				}
			}
		}
	}
	// 如果没有找到升序的位置，说明整个数组是降序的，直接反转数组即可
	reverse(0, len(nums)-1)
	return
}

func main() {
	nums := []int{3, 2, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}
