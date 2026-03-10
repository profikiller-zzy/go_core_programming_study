package main

func lengthOfLIS1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	var dp = make([]int, n)
	dp[0] = 1
	maxLen := 1
	for right := 1; right < n; right++ {
		var maxCurLen = 1
		for left := 0; left < right; left++ {
			if nums[right] > nums[left] {
				maxCurLen = max(dp[left]+1, maxCurLen)
			}
		}
		dp[right] = maxCurLen
		if maxLen < maxCurLen {
			maxLen = maxCurLen
		}
	}
	return maxLen
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// tails[i] 表示长度为 i+1 的递增子序列的最小结尾值
	tails := make([]int, 0, n)

	for _, x := range nums {
		// 如果 tails 为空，或者 x 大于 tails 的最后一个元素，直接追加
		if len(tails) == 0 || x > tails[len(tails)-1] {
			tails = append(tails, x)
		} else {
			// 否则，在 tails 中二分查找第一个 >= x 的位置，并替换它
			// 目的是让该长度的子序列结尾尽可能小
			left, right := 0, len(tails)-1
			for left <= right {
				mid := (left + right) / 2
				if tails[mid] >= x {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
			// left 即为要替换的位置
			tails[left] = x
		}
	}

	return len(tails)
}
