package main

func jump(nums []int) int {
	var (
		steps int
		start int
		end   int
	)
	for end < len(nums)-1 {
		steps++
		maxDis := end
		for index := start; index <= end; index++ {
			if index+nums[index] > maxDis {
				maxDis = index + nums[index]
			}
		}
		start = end + 1
		end = maxDis
	}
	return steps
}
