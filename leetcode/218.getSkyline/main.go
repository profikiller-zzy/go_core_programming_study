package main

import (
	"container/heap"
	"sort"
)

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type PriorityQueue struct {
	heap      maxHeap
	deleteCnt map[int]int
}

func NewPriorityQueue() *PriorityQueue {
	queue := &PriorityQueue{
		heap:      maxHeap{},
		deleteCnt: make(map[int]int),
	}
	heap.Init(&queue.heap)
	return queue
}

func (pq *PriorityQueue) Push(h int) {
	heap.Push(&pq.heap, h)
}

func (pq *PriorityQueue) Delete(h int) {
	pq.deleteCnt[h]++
}

func (pq *PriorityQueue) Peek() int {
	for pq.heap.Len() > 0 {
		top := pq.heap[0]
		if _, ok := pq.deleteCnt[top]; ok { // 检查当前高度是否需要弹出大根堆
			heap.Pop(&pq.heap)
			pq.deleteCnt[top]--
			if pq.deleteCnt[top] == 0 {
				delete(pq.deleteCnt, top)
			}
		} else {
			break
		}
	}
	if pq.heap.Len() == 0 {
		return 0
	}
	return pq.heap[0]
}

func getSkyline(buildings [][]int) [][]int {
	events := make([][2]int, 0, len(buildings)*2)
	for _, building := range buildings {
		left, right, height := building[0], building[1], building[2]
		events = append(events, [2]int{left, -height})
		events = append(events, [2]int{right, height})
	}
	sort.Slice(events, func(i, j int) bool {
		if events[i][0] != events[j][0] {
			return events[i][0] < events[j][0]
		}
		return events[i][1] < events[j][1]
	})
	pq := NewPriorityQueue()
	pq.Push(0)
	prev := 0
	result := make([][]int, 0)
	for _, event := range events {
		x, h := event[0], event[1]
		if h < 0 {
			pq.Push(-h)
		} else {
			pq.Delete(h)
		}

		top := pq.Peek()
		if top != prev {
			result = append(result, []int{x, top})
			prev = top
		}
	}
	return result
}
