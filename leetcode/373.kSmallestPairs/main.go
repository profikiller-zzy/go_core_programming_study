package main

import (
	"container/heap"
	"fmt"
)

type pair struct {
	sum            int
	index1, index2 int
}

type pairsHeap []pair

func (h pairsHeap) Len() int {
	return len(h)
}
func (h pairsHeap) Less(i, j int) bool {
	return h[i].sum < h[j].sum
}
func (h pairsHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *pairsHeap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}
func (h *pairsHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var res [][]int
	res = make([][]int, 0, k)

	visited := make(map[[2]int]bool)
	pHeap := &pairsHeap{}
	heap.Init(pHeap)

	// 初始化堆，只放 nums1[0] + nums2[0..min(k,len(nums2))] 的组合
	heap.Push(pHeap, pair{
		sum:    nums1[0] + nums2[0],
		index1: 0,
		index2: 0,
	})
	visited[[2]int{0, 0}] = true

	for k > 0 && pHeap.Len() > 0 {
		minPair := heap.Pop(pHeap).(pair)
		index1 := minPair.index1
		index2 := minPair.index2
		res = append(res, []int{nums1[index1], nums2[index2]})
		k--
		if index1+1 < len(nums1) && !visited[[2]int{index1 + 1, index2}] {
			heap.Push(pHeap, pair{
				sum:    nums1[index1+1] + nums2[index2],
				index1: index1 + 1,
				index2: index2,
			})
			visited[[2]int{index1 + 1, index2}] = true
		}
		if index2+1 < len(nums2) && !visited[[2]int{index1, index2 + 1}] {
			heap.Push(pHeap, pair{
				sum:    nums1[index1] + nums2[index2+1],
				index1: index1,
				index2: index2 + 1,
			})
			visited[[2]int{index1, index2 + 1}] = true
		}
	}
	return res
}

func main() {
	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}
	fmt.Println(kSmallestPairs(nums1, nums2, 3)) // [[1,2],[1,4],[1,6]]
}
