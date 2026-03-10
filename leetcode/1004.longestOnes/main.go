package main

func longestOnes(nums []int, k int) int {
	var left, right int
	for right < len(nums) {
		if nums[right] == 0 {
			k--
		}
		right++

		// k < 0 代表现在翻转已经用完了，窗口内的0多于K，第一次k=-1时窗口一定是由于当前循环扩展右边届遇到0导致的
		// 只要k < 0就需要收缩左边界，k的值可能会取-1 -2等等，但是没有关系，此时right - left依旧等于历史最大合法窗口的长度
		// 因为每次循环右边界是一定扩展的，只要k < 0，左边界收缩，窗口大小是不变的
		// 窗口代表的范围是[left, right)
		if k < 0 {
			if nums[left] == 0 {
				k++
			}
			left++
		}
	}
	return right - left
}
