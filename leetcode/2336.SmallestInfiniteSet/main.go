package main

import (
	"container/heap"
)

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type SmallestInfiniteSet struct {
	offset int     // 大于等于offset的所有整数都在集合中
	heap   minHeap // 只存储小于offset的正整数
	inHeap map[int]bool
}

func Constructor() SmallestInfiniteSet {
	set := SmallestInfiniteSet{
		offset: 1,
		heap:   make(minHeap, 0),
		inHeap: make(map[int]bool),
	}
	heap.Init(&set.heap)
	return set
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if this.heap.Len() == 0 {
		this.offset++
		return this.offset - 1
	}
	x := heap.Pop(&this.heap).(int)
	this.inHeap[x] = false
	return x
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num >= this.offset {
		return
	}
	if !this.inHeap[num] {
		heap.Push(&this.heap, num)
		this.inHeap[num] = true
	}
}
