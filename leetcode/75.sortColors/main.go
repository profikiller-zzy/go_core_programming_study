package main

import "fmt"

func sortColors(nums []int) {
	// low 下一个应该放 0 的位置； mid 正在处理的位置； high 下一个应该放 2 的位置
	low, mid, high := 0, 0, len(nums)-1

	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}
}

func main() {
	nums := []int{1, 0, 1}
	sortColors(nums)
	fmt.Println(nums)
}
