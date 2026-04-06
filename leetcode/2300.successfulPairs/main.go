package main

import (
	"fmt"
	"sort"
)

func successfulPairs1(spells []int, potions []int, success int64) []int {
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

func successfulPairs(spells []int, potions []int, success int64) []int {
	result := make([]int, len(spells))
	sort.Slice(potions, func(i, j int) bool { return potions[i] < potions[j] })
	for i := 0; i < len(spells); i++ {
		spell := spells[i]
		left, right := 0, len(potions)-1
		for left <= right {
			mid := (left + right) / 2
			if int64(potions[mid])*int64(spell) < success {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		result[i] = len(potions) - left
	}
	return result
}

func main() {
	spells := []int{5, 1, 3}
	potions := []int{1, 2, 3, 4, 5}
	result := successfulPairs(spells, potions, 7)
	fmt.Println(result)
}
