package main

func jump(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	jumps := 0
	curEnd := 0  // 当前这一步能到达的最远边界
	nextEnd := 0 // 下一步能到达的最远边界

	for i := 0; i < n-1; i++ {
		if i+nums[i] > nextEnd {
			nextEnd = i + nums[i]
		}
		if i == curEnd {
			jumps++
			curEnd = nextEnd
			if curEnd >= n-1 {
				break
			}
		}
	}

	return jumps
}
