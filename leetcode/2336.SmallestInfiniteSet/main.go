package main

type minHeap struct {
	data []int
}

func newMinHeap() *minHeap {
	return &minHeap{
		data: make([]int, 0),
	}
}

func (h minHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h minHeap) Len() int {
	return len(h.data)
}

func (h *minHeap) UpAdjust(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.data[parentIndex] > h.data[index] {
			h.Swap(parentIndex, index)
			index = parentIndex
		} else {
			break
		}
	}
}

func (h *minHeap) DownAdjust(index int) {
	maxIndex := len(h.data) - 1
	for {
		leftChild, rightChild, minIndex := index*2+1, index*2+2, index
		if leftChild <= maxIndex && h.data[leftChild] < h.data[minIndex] {
			minIndex = leftChild
		}
		if rightChild <= maxIndex && h.data[rightChild] < h.data[minIndex] {
			minIndex = rightChild
		}
		if minIndex == index {
			break
		}
		h.Swap(minIndex, index)
		index = minIndex
	}
}

func (h *minHeap) Insert(value int) {
	h.data = append(h.data, value)
	h.UpAdjust(len(h.data) - 1)
}

func (h *minHeap) RemoveMin() int {
	minValue := h.data[0]
	h.Swap(0, len(h.data)-1)
	h.data = h.data[:len(h.data)-1]
	h.DownAdjust(0)
	return minValue
}

type SmallestInfiniteSet struct {
	h      *minHeap
	minNum int
	inHeap map[int]bool
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		h:      newMinHeap(),
		minNum: 1,
		inHeap: make(map[int]bool),
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if this.h.Len() > 0 {
		minValue := this.h.RemoveMin()
		delete(this.inHeap, minValue)
		return minValue
	}
	this.minNum++
	return this.minNum - 1
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num >= this.minNum {
		return
	}
	if this.inHeap[num] {
		return
	}
	this.inHeap[num] = true
	this.h.Insert(num)
}
