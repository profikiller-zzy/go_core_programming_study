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

func kSmallestPairs1(nums1 []int, nums2 []int, k int) [][]int {
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

type minHeap struct {
	data []pair
}

func newMinHeap() *minHeap { // 建议返回指针
	return &minHeap{
		data: make([]pair, 0),
	}
}

func (h *minHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *minHeap) Len() int {
	return len(h.data)
}

func (h *minHeap) upAdjust(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		// 小顶堆：父节点应该小于子节点
		if h.data[parentIndex].sum > h.data[index].sum {
			h.Swap(index, parentIndex)
			index = parentIndex
		} else {
			break
		}
	}
}

func (h *minHeap) downAdjust(index int) {
	n := h.Len()
	for {
		leftChild, rightChild := index*2+1, index*2+2
		minIndex := index
		if leftChild < n && h.data[leftChild].sum < h.data[minIndex].sum {
			minIndex = leftChild
		}
		if rightChild < n && h.data[rightChild].sum < h.data[minIndex].sum {
			minIndex = rightChild
		}
		if minIndex == index {
			return
		}
		h.Swap(index, minIndex)
		index = minIndex
	}
}

func (h *minHeap) Push(p pair) {
	h.data = append(h.data, p)
	h.upAdjust(h.Len() - 1)
}

func (h *minHeap) Pop() pair {
	top := h.data[0]
	h.Swap(0, h.Len()-1)
	h.data = h.data[:h.Len()-1]
	if h.Len() > 0 {
		h.downAdjust(0)
	}
	return top
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	isVisited := make(map[[2]int]bool)
	h := newMinHeap()

	h.Push(pair{nums1[0] + nums2[0], 0, 0})
	isVisited[[2]int{0, 0}] = true

	var result [][]int

	for k > 0 && h.Len() > 0 {
		// 使用封装好的 Pop
		minPair := h.Pop()

		index1, index2 := minPair.index1, minPair.index2
		result = append(result, []int{nums1[index1], nums2[index2]})
		k--

		// 向右扩展 (nums1 的下一个元素)
		if index1+1 < len(nums1) && !isVisited[[2]int{index1 + 1, index2}] {
			h.Push(pair{nums1[index1+1] + nums2[index2], index1 + 1, index2})
			isVisited[[2]int{index1 + 1, index2}] = true
		}

		// 向下扩展 (nums2 的下一个元素)
		if index2+1 < len(nums2) && !isVisited[[2]int{index1, index2 + 1}] {
			h.Push(pair{nums1[index1] + nums2[index2+1], index1, index2 + 1})
			isVisited[[2]int{index1, index2 + 1}] = true
		}
	}
	return result
}

func main() {
	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}
	fmt.Println(kSmallestPairs(nums1, nums2, 3)) // [[1,2],[1,4],[1,6]]
}
