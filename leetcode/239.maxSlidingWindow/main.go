package main

import (
	"container/heap"
	"fmt"
	"sort"
)

var a []int

type hp struct{ sort.IntSlice }

func (h *hp) Less(i, j int) bool { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func maxSlidingWindow1(nums []int, k int) []int {
	a = nums
	q := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q)

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		ans = append(ans, nums[q.IntSlice[0]])
	}
	return ans
}

// maxSlidingWindow 使用单调队列实现窗口最大值
func maxSlidingWindow(nums []int, k int) []int {
	var queue = make([]int, 0)
	push := func(index int) {
		// 如果len(queue) > 0，代表队列中有之前的元素，index为当前元素的索引
		// 加入新元素索引，同时维持队列的递减性，这样可以保证队头始终是当前窗口中最大元素的索引
		for len(queue) > 0 && nums[index] > nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, index)
	}

	var res = make([]int, 1, len(nums)-k+1)
	for index := 0; index < k; index++ {
		push(index)
	}
	res[0] = nums[queue[0]]

	for index := k; index < len(nums); index++ {
		push(index)
		for queue[0] <= index-k { // 将不在窗口中的索引（不在窗口范围内）移出
			queue = queue[1:]
		}
		res = append(res, nums[queue[0]])
	}
	return res
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	result := maxSlidingWindow(nums, k)
	fmt.Println(result)
}
