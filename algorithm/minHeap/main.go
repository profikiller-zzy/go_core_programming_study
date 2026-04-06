package main

import (
	"errors"
	"fmt"
)

// MinHeap 最小堆结构体
type MinHeap struct {
	data []int
}

// NewMinHeap 创建新的最小堆
func NewMinHeap() *MinHeap {
	return &MinHeap{
		data: make([]int, 0),
	}
}

// parent 获取父节点索引
func parent(i int) int {
	return (i - 1) / 2
}

// left 获取左子节点索引
func left(i int) int {
	return 2*i + 1
}

// right 获取右子节点索引
func right(i int) int {
	return 2*i + 2
}

// Insert 插入元素到堆中
func (h *MinHeap) Insert(value int) {
	h.data = append(h.data, value)
	h.heapifyUp(len(h.data) - 1)
}

// heapifyUp 向上调整堆
func (h *MinHeap) heapifyUp(index int) {
	for index > 0 {
		parentIndex := parent(index)
		if h.data[index] < h.data[parentIndex] {
			h.swap(index, parentIndex)
			index = parentIndex
		} else {
			break
		}
	}
}

// ExtractMin 提取并移除最小元素
func (h *MinHeap) ExtractMin() (int, error) {
	if len(h.data) == 0 {
		return 0, errors.New("heap is empty")
	}

	min := h.data[0]

	// 将最后一个元素移到根节点
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]

	// 向下调整
	if len(h.data) > 0 {
		h.heapifyDown(0)
	}

	return min, nil
}

// heapifyDown 向下调整堆
func (h *MinHeap) heapifyDown(index int) {
	lastIndex := len(h.data) - 1

	for {
		leftIndex := left(index)
		rightIndex := right(index)
		smallest := index

		// 找到最小的子节点
		if leftIndex <= lastIndex && h.data[leftIndex] < h.data[smallest] {
			smallest = leftIndex
		}

		if rightIndex <= lastIndex && h.data[rightIndex] < h.data[smallest] {
			smallest = rightIndex
		}

		// 如果当前节点已经是最小的，停止调整
		if smallest == index {
			break
		}

		// 交换并继续向下调整
		h.swap(index, smallest)
		index = smallest
	}
}

// swap 交换两个元素
func (h *MinHeap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// Peek 查看最小元素但不移除
func (h *MinHeap) Peek() (int, error) {
	if len(h.data) == 0 {
		return 0, errors.New("heap is empty")
	}
	return h.data[0], nil
}

// Size 返回堆的大小
func (h *MinHeap) Size() int {
	return len(h.data)
}

// IsEmpty 检查堆是否为空
func (h *MinHeap) IsEmpty() bool {
	return len(h.data) == 0
}

// BuildHeap 从切片构建堆
func (h *MinHeap) BuildHeap(arr []int) {
	h.data = make([]int, len(arr))
	copy(h.data, arr)

	// 从最后一个非叶子节点开始向下调整
	for i := len(h.data)/2 - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}
}

// 示例使用
func main() {
	// 创建最小堆
	heap := NewMinHeap()

	// 插入元素
	heap.Insert(10)
	heap.Insert(5)
	heap.Insert(20)
	heap.Insert(3)
	heap.Insert(15)

	fmt.Println("堆的大小:", heap.Size())

	// 提取最小元素
	for !heap.IsEmpty() {
		min, _ := heap.ExtractMin()
		fmt.Printf("提取最小值: %d\n", min)
	}

	fmt.Println("\n从数组构建堆:")
	arr := []int{9, 5, 7, 1, 3, 8}
	heap.BuildHeap(arr)
	fmt.Println("原始数组:", arr)
	fmt.Println("堆化后:", heap.data)

	// 查看最小元素
	if min, err := heap.Peek(); err == nil {
		fmt.Printf("堆顶元素: %d\n", min)
	}
}
