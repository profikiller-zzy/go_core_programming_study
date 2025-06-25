package main

import "strconv"

func evalRPN(tokens []string) int {
	nums := make([]int, 0, len(tokens))
	var curVal int
	for index := 0; index < len(tokens); index++ {
		switch tokens[index] {
		case "+":
			if len(nums) < 2 {
				return 0 // 错误处理
			}
			curVal = nums[len(nums)-2] + nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			nums[len(nums)-1] = curVal
		case "-":
			if len(nums) < 2 {
				return 0 // 错误处理
			}
			curVal = nums[len(nums)-2] - nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			nums[len(nums)-1] = curVal
		case "*":
			if len(nums) < 2 {
				return 0 // 错误处理
			}
			curVal = nums[len(nums)-2] * nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			nums[len(nums)-1] = curVal
		case "/":
			if len(nums) < 2 {
				return 0 // 错误处理
			}
			curVal = nums[len(nums)-2] / nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			nums[len(nums)-1] = curVal
		default:
			// 处理数字
			num, _ := strconv.ParseInt(tokens[index], 10, 64)
			nums = append(nums, int(num))
		}
	}
	return nums[0] // 最后一个元素即为结果
}
