package main

import "fmt"

// https://leetcode.cn/problems/product-of-array-except-self/description

func productExceptSelf(nums []int) []int {
	// 先求出前缀积和后缀积
	prefix, suffix := make([]int, len(nums)), make([]int, len(nums))
	prefix[0] = nums[0]
	suffix[len(nums)-1] = nums[len(nums)-1]
	for index := 1; index < len(nums); index++ {
		prefix[index] = prefix[index-1] * nums[index]
	}
	for index := len(nums) - 2; index >= 0; index-- {
		suffix[index] = suffix[index+1] * nums[index]
	}

	// 那么ans[i] = prefix[i-1] * suffix[i+1]
	ans := make([]int, len(nums))
	ans[0] = suffix[1]
	ans[len(nums)-1] = prefix[len(nums)-2]
	for index := 1; index < len(nums)-1; index++ {
		ans[index] = prefix[index-1] * suffix[index+1]
	}
	return ans
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
