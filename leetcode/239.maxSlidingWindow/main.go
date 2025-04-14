package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type winnerTree struct {
	players []int // 存储选手的值
	tree    []int // 存储树的节点
	k       int   // 选手的数量
}

func newWinnerTree(k int, players []int) *winnerTree {
	wt := &winnerTree{
		players: make([]int, k),
		tree:    make([]int, 2*k-1),
		k:       k,
	}
	wt.init(players)
	return wt
}

func (wt *winnerTree) init(players []int) {
	wt.players = players
	wt.buildTree()
}

func (wt *winnerTree) buildTree() {
	for i := 0; i < wt.k; i++ {
		wt.tree[wt.k-1+i] = i
	}
	for i := wt.k - 2; i >= 0; i-- {
		if wt.players[wt.tree[2*i+1]] > wt.players[wt.tree[2*i+2]] {
			wt.tree[i] = wt.tree[2*i+1]
		} else {
			wt.tree[i] = wt.tree[2*i+2]
		}
	}
}

func (wt *winnerTree) returnWinner() int {
	return wt.players[wt.tree[0]]
}

func (wt *winnerTree) update(index int, value int) {
	wt.players[index] = value
	treeIndex := wt.k - 1 + index
	for treeIndex > 0 {
		parent := (treeIndex - 1) / 2
		if wt.players[wt.tree[2*parent+1]] > wt.players[wt.tree[2*parent+2]] {
			wt.tree[parent] = wt.tree[2*parent+1]
		} else {
			wt.tree[parent] = wt.tree[2*parent+2]
		}
		treeIndex = parent
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	winnerTree := newWinnerTree(k, nums[:k])
	var result []int
	for i := k; i < len(nums); i++ {
		result = append(result, winnerTree.returnWinner())
		winnerTree.update(i%k, nums[i])
	}
	result = append(result, winnerTree.returnWinner())
	return result
}

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

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	result := maxSlidingWindow(nums, k)
	fmt.Println(result)
}
