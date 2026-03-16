package main

import (
	"container/heap"
	"fmt"
)

type pair struct {
	index, cost int
}

type minHeap []pair

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Less(i, j int) bool {
	// 关键，如果两个工人的cost相同，则下标小的排在前面
	if h[i].cost == h[j].cost {
		return h[i].index < h[j].index
	}
	return h[i].cost < h[j].cost
}
func (h minHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func totalCost(costs []int, k int, candidates int) int64 {
	n := len(costs)
	h := &minHeap{}
	heap.Init(h)
	left, right := candidates, n-candidates-1
	for i := 0; i < n; i++ {
		if i < left || i > right {
			heap.Push(h, pair{i, costs[i]})
		}
	}

	var tCost int64
	for i := 0; i < k; i++ {
		p := heap.Pop(h).(pair)
		tCost += int64(p.cost)
		index := p.index
		if left <= right {
			if index < left { // 工人来自左边
				heap.Push(h, pair{left, costs[left]})
				left++
			} else {
				heap.Push(h, pair{right, costs[right]})
				right--
			}
		}
	}
	return tCost
}

func main() {
	costs := []int{2, 2, 2, 2, 2, 2, 1, 4, 5, 5, 5, 5, 5, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
	k, candidates := 7, 3
	fmt.Println(totalCost(costs, k, candidates))
}
