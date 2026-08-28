package main

import (
	"fmt"
)

type HeapInterface interface {
	Push(value int)
	Pop() int
	Peek() int
	Len() int
}

type Heap struct {
	data []int
	// less(a, b) 为 true 表示 a 的优先级高于 b
	less func(a, b int) bool
}

func NewHeap(isMinHeap bool) HeapInterface {
	less := func(a, b int) bool {
		if isMinHeap {
			return a < b
		}
		return a > b
	}

	return &Heap{
		data: make([]int, 0),
		less: less,
	}
}

func (h *Heap) Len() int {
	return len(h.data)
}

func (h *Heap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap) UpAdjust(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if !h.less(h.data[index], h.data[parentIndex]) {
			break
		}
		h.swap(index, parentIndex)
		index = parentIndex
	}
}

func (h *Heap) DownAdjust(index int) {
	n := h.Len()

	for {
		leftChild, rightChild := 2*index+1, 2*index+2
		bestIndex := index
		if leftChild < n && h.less(h.data[leftChild], h.data[bestIndex]) {
			bestIndex = leftChild
		}
		if rightChild < n && h.less(h.data[rightChild], h.data[bestIndex]) {
			bestIndex = rightChild
		}
		if bestIndex == index {
			return
		}
		h.swap(index, bestIndex)
		index = bestIndex
	}
}

func (h *Heap) Push(value int) {
	h.data = append(h.data, value)
	h.UpAdjust(h.Len() - 1)
}

func (h *Heap) Pop() int {
	if h.Len() == 0 {
		return -1
	}

	top := h.data[0]
	h.swap(0, h.Len()-1)
	h.data = h.data[:h.Len()-1]
	h.DownAdjust(0)
	return top
}

func (h *Heap) Peek() int {
	if h.Len() == 0 {
		return -1
	}
	return h.data[0]
}

type MedianFinder struct {
	minHeap, maxHeap HeapInterface
}

func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: NewHeap(true),
		maxHeap: NewHeap(false),
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.minHeap.Len() == 0 || num >= this.minHeap.Peek() {
		this.minHeap.Push(num)
	} else {
		this.maxHeap.Push(num)
	}

	if this.minHeap.Len() > this.maxHeap.Len()+1 {
		this.maxHeap.Push(this.minHeap.Pop())
	} else if this.maxHeap.Len() > this.minHeap.Len() {
		this.minHeap.Push(this.maxHeap.Pop())
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if (this.minHeap.Len()+this.maxHeap.Len())%2 == 0 {
		return float64(this.minHeap.Peek()+this.maxHeap.Peek()) / 2.0
	}
	return float64(this.minHeap.Peek())
}
func main() {
	medianFinder := Constructor()
	medianFinder.AddNum(1)
	medianFinder.AddNum(2)
	medianFinder.AddNum(3)
	fmt.Println(medianFinder.FindMedian()) // 应输出 2
}
