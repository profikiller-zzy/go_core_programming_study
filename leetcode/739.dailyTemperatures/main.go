package main

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0) // 构建一个递减栈
	res := make([]int, len(temperatures))
	for index := 0; index < len(temperatures); index++ {
		for len(stack) > 0 && temperatures[index] > temperatures[stack[len(stack)-1]] {
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[prev] = index - prev
		}
		stack = append(stack, index)
	}
	return res
}
