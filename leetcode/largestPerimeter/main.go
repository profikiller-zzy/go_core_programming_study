package main

import "sort"

func main() {

}

func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	for l1 := len(nums) - 1; l1 >= 2; l1-- {
		for l2 := l1 - 1; l2 >= 1; l2-- {
			for l3 := l2 - 1; l3 >= 0; l3-- {
				if nums[l2]+nums[l3] > nums[l1] {
					return nums[l2] + nums[l3] + nums[l1]
				} else {
					break
				}
			}
		}
	}
	return 0
}
