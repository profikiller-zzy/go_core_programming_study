package main

type Stack struct {
	len  int
	data []int
}

func newStack() *Stack {
	return &Stack{
		len:  0,
		data: make([]int, 0),
	}
}

func (s *Stack) push(c int) {
	s.data = append(s.data, c)
	s.len++
}

func (s *Stack) pop() {
	if s.len == 0 {
		return
	}
	s.data = s.data[:s.len-1]
	s.len--
}

func (s *Stack) top() int {
	if s.len == 0 {
		return -1
	}
	return s.data[s.len-1]
}

func (s *Stack) isEmpty() bool {
	return s.len == 0
}

// dailyTemperatures 使用递减栈解决每日温度问题
func dailyTemperatures(temperatures []int) []int {
	var result []int
	result = make([]int, len(temperatures))
	stack := newStack()
	for index := 0; index < len(temperatures); index++ {
		for !stack.isEmpty() && temperatures[index] > temperatures[stack.top()] {
			top := stack.top()
			result[top] = index - top
			stack.pop()
		}
		stack.push(index)
	}
	for !stack.isEmpty() {
		top := stack.top()
		result[top] = 0
		stack.pop()
	}
	return result
}
