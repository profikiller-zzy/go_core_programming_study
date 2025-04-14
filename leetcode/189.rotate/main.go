package main

// https://leetcode.cn/problems/rotate-array/description

func rotate(nums []int, k int) {
	k = k % len(nums)
	if len(nums) <= 1 {
		return
	}
	
	reverse := func(start, end int) {
		for start < end {
			nums[start], nums[end] = nums[end], nums[start]
			start++
			end--
		}
	}
	reverse(0, len(nums)-1-k)
	reverse(len(nums)-k, len(nums)-1)
	reverse(0, len(nums)-1)
}

func main() {

}
