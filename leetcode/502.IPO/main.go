package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}
func (h intHeap) Less(i, j int) bool { // 建立大根堆
	return h[i] > h[j]
}
func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)
	type pair struct {
		profit  int
		capital int
	}
	pairs := make([]pair, n)
	for index := 0; index < n; index++ {
		pairs[index] = pair{profits[index], capital[index]}
	}
	sort.Slice(pairs, func(i, j int) bool { // 将它们按照所需资本大小进行排序
		return pairs[i].capital < pairs[j].capital
	})
	maxHeap := &intHeap{}
	heap.Init(maxHeap)
	// 该算法的核心思想就是每次从当前可用的项目中选择利润最大的项目进行投资，直到达到k次投资或没有可用项目为止。
	// 那么当前可用的项目就是所需的资本小于等于当前所持有的资本的项目
	for index := 0; k > 0; k-- {
		for index < n && pairs[index].capital <= w {
			heap.Push(maxHeap, pairs[index].profit)
			index++
		}
		if maxHeap.Len() == 0 {
			break
		} else {
			w += heap.Pop(maxHeap).(int)
		}
	}
	return w
}

func main() {
	var (
		k, w             int
		profits, capital []int
	)
	k = 3
	w = 0
	profits = []int{1, 2, 3}
	capital = []int{0, 1, 1}
	result := findMaximizedCapital(k, w, profits, capital)
	fmt.Println(result)
}
