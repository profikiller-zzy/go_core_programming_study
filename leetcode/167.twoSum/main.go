package main

import "fmt"

func twoSum1(numbers []int, target int) []int {
	if len(numbers) == 2 { // 特例处理
		return []int{1, 2}
	}

	var index1 int
	value1, value2 := numbers[0], target-numbers[0]
	for index1 = 0; index1 <= len(numbers)-2; index1++ { // 加数1
		if numbers[index1] == value1 && index1 > 0 { // 如果当前加数和上一个加数相同，跳过
			continue
		} else {
			value1 = numbers[index1]
		}
		value2 = target - value1

		// 二分查找第二个加数
		left, right := index1+1, len(numbers)-1
		for left <= right {
			mid := (left + right) / 2
			if numbers[mid] > value2 {
				right = mid - 1
			} else if numbers[mid] < value2 {
				left = mid + 1
			} else {
				return []int{index1 + 1, mid + 1}
			}
		}
	}
	return nil
}

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}

func main() {
	nums := []int{-1, 0}
	target := -1
	fmt.Println(twoSum(nums, target))
}
