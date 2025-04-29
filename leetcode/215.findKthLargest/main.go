package main

import (
	"container/heap"
	"fmt"
)

// IntHeap 定义一个整型堆
type IntHeap []int

// 实现 heap.Interface 接口
func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] } // 小的元素优先
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	if k > len(nums) {
		return -1
	}
	// 创建一个小顶堆
	minHeap := &IntHeap{}
	heap.Init(minHeap) // 初始化堆
	for i := 0; i < len(nums); i++ {
		if minHeap.Len() < k {
			heap.Push(minHeap, nums[i])
		} else {
			if nums[i] > (*minHeap)[0] {
				heap.Pop(minHeap) // 弹出堆顶元素
				heap.Push(minHeap, nums[i])
			}
		}
	}
	return (*minHeap)[0] // 堆顶元素就是第k大的元素
}

func main() {
	nums := []int{2, 1}
	fmt.Println(findKthLargest(nums, 2)) // 输出 1
}
