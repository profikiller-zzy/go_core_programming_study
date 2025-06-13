package main

func removeDuplicates(nums []int) int {
	var slow, fast int
	for ; fast < len(nums); fast++ {
		if fast == 0 {
			continue
		}
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}
