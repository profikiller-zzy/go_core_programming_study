package main

import "container/heap"

type frequency struct {
	num   int
	times int
}

// 首先实现最小堆
type minHeap []frequency

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i].times < h[j].times
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(frequency))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	frequent := make(map[int]int)
	for _, num := range nums {
		frequent[num]++
	}

	mHeap := minHeap{}
	heap.Init(&mHeap)
	// 使用最小堆来存储频率
	for num, freq := range frequent {
		if mHeap.Len() < k {
			heap.Push(&mHeap, frequency{num, freq})
		} else {
			if freq > mHeap[0].times {
				heap.Pop(&mHeap)
				heap.Push(&mHeap, frequency{num, freq})
			}
		}
	}
	var res = make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = heap.Pop(&mHeap).(frequency).num
	}
	return res
}
