package main

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	var slow, fast, count int
	for ; fast < len(nums); fast++ {
		if fast == 0 {
			count++
			continue
		}
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
			count = 1
		} else {
			if count < 2 {
				count++
				slow++
				nums[slow] = nums[fast]
			} else {
				continue // Skip duplicates
			}
		}
	}
	return slow + 1
}
