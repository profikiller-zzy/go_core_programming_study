package main

func removeElement(nums []int, val int) int {
	var left, right int
	for right < len(nums) {
		if nums[right] == val {
			right++
		} else {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right++
		}
	}
	return left
}
