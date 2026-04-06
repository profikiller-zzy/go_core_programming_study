package main

// leetcode 26. 删除有序数组中的重复项
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/

func removeDuplicates(nums []int) int {
	var curIndex, newIndex int
	for ; curIndex < len(nums); curIndex++ {
		if curIndex == 0 {
			continue
		}
		if nums[newIndex] != nums[curIndex] {
			newIndex++
			nums[newIndex] = nums[curIndex]
		}
	}
	return newIndex + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	println(removeDuplicates(nums))
}
