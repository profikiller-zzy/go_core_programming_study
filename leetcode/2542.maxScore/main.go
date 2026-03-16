package main

import (
	"container/heap"
	"sort"
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

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	n := len(nums1)
	pairs := make([][2]int, n)
	for i := 0; i < n; i++ {
		pairs[i] = [2]int{nums1[i], nums2[i]}
	}
	sort.Slice(pairs, func(i, j int) bool { // 按照nums2中的元素降序排列
		return pairs[i][1] > pairs[j][1]
	})

	h := &minHeap{}
	heap.Init(h)
	var mScore, curScore, curSum int64
	for i := 0; i < n; i++ {
		num1, num2 := pairs[i][0], pairs[i][1]
		heap.Push(h, num1)
		curSum += int64(num1)
		if h.Len() > k { // 堆满了，则弹出最小元素
			top := heap.Pop(h).(int)
			curSum -= int64(top)
		}

		if h.Len() == k {
			// 计算当前分数
			curScore = curSum * int64(num2)
			if curScore > mScore {
				mScore = curScore
			}
		}
	}
	return mScore
}
