package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// 小顶堆，存放大于等于中位数的那一部分
type minHeap struct{ sort.IntSlice }

func (h *minHeap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}
func (h *minHeap) Pop() interface{} {
	old := h.IntSlice
	n := len(old)
	x := old[n-1]
	h.IntSlice = old[0 : n-1]
	return x
}

// 大顶堆，存放小于中位数的数
type maxHeap struct{ sort.IntSlice }

func (h *maxHeap) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // 实现大顶堆
func (h *maxHeap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}
func (h *maxHeap) Pop() interface{} {
	old := h.IntSlice
	n := len(old)
	x := old[n-1]
	h.IntSlice = old[0 : n-1]
	return x
}

type MedianFinder struct {
	sHeap *minHeap // 小顶堆，存储大于等于中位数的数
	bHeap *maxHeap // 大顶堆，存储小于中位数的数
}

func Constructor() MedianFinder {
	return MedianFinder{
		sHeap: &minHeap{},
		bHeap: &maxHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	// 小顶堆用于存放大于等于中位数的数，而大顶堆用来存放小于中位数的数
	if this.sHeap.Len() == 0 || num >= this.sHeap.IntSlice[0] {
		heap.Push(this.sHeap, num)
	} else {
		heap.Push(this.bHeap, num)
	}

	// 调整两个堆中元素的数量，使得小顶堆的元素个数不超过大顶堆的元素个数加1
	if this.sHeap.Len() > this.bHeap.Len()+1 {
		heap.Push(this.bHeap, heap.Pop(this.sHeap).(int))
	} else if this.bHeap.Len() > this.sHeap.Len() {
		heap.Push(this.sHeap, heap.Pop(this.bHeap).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if (this.bHeap.Len()+this.sHeap.Len())%2 == 0 {
		return float64(this.bHeap.IntSlice[0]+this.sHeap.IntSlice[0]) / 2.0
	}
	return float64(this.sHeap.IntSlice[0])
}

func main() {
	medianFinder := Constructor()
	medianFinder.AddNum(1)
	medianFinder.AddNum(2)
	medianFinder.AddNum(3)
	fmt.Println(medianFinder.FindMedian()) // 应输出 2
}
