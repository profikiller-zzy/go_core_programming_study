package main

func findMin(nums []int) int {
	var (
		low  = 0
		high = len(nums) - 1
		mid  = 0
	)
	for low < high {
		mid = (low + high) / 2
		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return nums[low]
}
