package main

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	m, n := len(spells), len(potions)
	sort.Ints(potions)
	result := make([]int, m)

	for i, spell := range spells {
		target := (success + int64(spell) - 1) / int64(spell)

		// 标准二分查找第一个 >= target 的位置
		left, right := 0, n-1
		for left <= right {
			mid := left + (right-left)/2
			if target >= int64(potions[mid]) {
				// right指向最后一个小于target的元素索引
				// 而left一定为right+1，所以left 指向第一个 >= target 的位置（可能为 n）
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		result[i] = n - left
	}
	return result
}
