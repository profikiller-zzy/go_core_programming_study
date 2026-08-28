package main

type MinStack struct {
	stack  []int
	minVal int
}

func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
	}
}

func (this *MinStack) Push(value int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, 0)
		this.minVal = value
	} else {
		diff := value - this.minVal
		this.stack = append(this.stack, diff)
		if diff < 0 {
			this.minVal = value
		}
	}
}

func (this *MinStack) Pop() {
	top := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if top < 0 {
		this.minVal = this.minVal - top
	}
}

func (this *MinStack) Top() int {
	top := this.stack[len(this.stack)-1]
	if top < 0 {
		return this.minVal
	}
	return this.minVal + top
}

func (this *MinStack) GetMin() int {
	return this.minVal
}
