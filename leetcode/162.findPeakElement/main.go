package main

import "math"

func findPeakElement(nums []int) int {
	n := len(nums)

	var getValue func(index int) int
	getValue = func(index int) int {
		if index == -1 || index == n {
			return math.MinInt64
		}
		return nums[index]
	}

	// 数组的边界也可以是峰值，所以二分查找最后总是能找到一个峰值，为什么呢？
	left, right := 0, n-1
	for {
		mid := (left + right) / 2
		if getValue(mid-1) < getValue(mid) && getValue(mid) > getValue(mid+1) {
			return mid
		} else if getValue(mid) < getValue(mid+1) { // 说明说明右边有上坡趋势 ⇒ 右侧一定有峰值 ⇒ 直接舍弃左边
			left = mid + 1
		} else { // 反之，直接舍弃右边
			right = mid - 1
		}
	}
}
