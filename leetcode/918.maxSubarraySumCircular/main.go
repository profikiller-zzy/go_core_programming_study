package main

func maxSubarraySumCircular(nums []int) int {
	total, maxSum, curMax := nums[0], nums[0], nums[0]
	minSum, curMin := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		curMax = max(curMax+num, num)
		maxSum = max(maxSum, curMax)

		curMin = min(curMin+num, num)
		minSum = min(minSum, curMin)

		total += num
	}

	// 全为负数的情况，total == minSum
	if maxSum > 0 {
		return max(maxSum, total-minSum)
	}
	// 如果说maxSum都小于等于零，则数组数的所有数字都小于等于零，这时候minSum等于所有数字之和，那么total-minSum一定为0
	// 直接返回这个total-minSum是会报错的，所以需要返回maxSum
	return maxSum
}

func main() {
	nums := []int{1, -2, 3, -2}
	result := maxSubarraySumCircular(nums)
	println(result) // Output: 3
}
