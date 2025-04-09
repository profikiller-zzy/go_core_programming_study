package main

// https://leetcode.cn/problems/min-stack/description/

// MinStack 使用两个栈实现最小栈
type MinStack struct {
	stack    []int
	stackMin []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		stackMin: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.stackMin) == 0 || val <= this.stackMin[len(this.stackMin)-1] {
		this.stackMin = append(this.stackMin, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	if this.stackMin[len(this.stackMin)-1] == this.stack[len(this.stack)-1] {
		this.stackMin = this.stackMin[:len(this.stackMin)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.stackMin[len(this.stackMin)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
