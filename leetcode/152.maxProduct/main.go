package main

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxDp, minDp := nums[0], nums[0]
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		tempMax := maxDp // 临时变量防止 minDp 更新前被覆盖

		// 以cur为结尾的最大值和最小值，即(当前值、前一个最大值*当前值、前一个最小值*当前值)的最小值或者最大值
		maxDp = max(cur, tempMax*cur, minDp*cur)
		minDp = min(cur, tempMax*cur, minDp*cur)

		result = max(result, maxDp)
	}

	return result
}

func main() {

}
