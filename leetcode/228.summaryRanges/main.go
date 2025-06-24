package main

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	res := make([]string, 0)
	for index := 0; index < len(nums); {
		start := index
		for index <= len(nums)-2 && nums[index+1] == nums[index]+1 {
			index++
		}
		end := index
		if start == end {
			res = append(res, fmt.Sprintf("%d", nums[start]))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[end]))
		}
		index++
	}
	return res
}

func main() {
	nums := []int{0, 1, 2, 4, 5, 7, 8}
	fmt.Println(summaryRanges(nums))
}
