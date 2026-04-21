package main

type minHeap struct {
	data []int
}

func (h minHeap) Len() int       { return len(h.data) }
func (h *minHeap) Swap(i, j int) { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *minHeap) UpAdjust(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.data[parentIndex] > h.data[index] {
			h.Swap(parentIndex, index)
			index = parentIndex
		} else {
			return
		}
	}
}

func (h *minHeap) DownAdjust(index int) {
	n := h.Len()
	for {
		left, right, minIndex := index*2+1, index*2+2, index
		if left < n && h.data[left] < h.data[minIndex] {
			minIndex = left
		}
		if right < n && h.data[right] < h.data[minIndex] {
			minIndex = right
		}
		if minIndex == index {
			return
		}
		h.Swap(minIndex, index)
		index = minIndex
	}
}

func (h *minHeap) Push(x int) {
	h.data = append(h.data, x)
	h.UpAdjust(h.Len() - 1)
}

func (h *minHeap) Pop() int {
	x := h.data[0]
	h.Swap(0, h.Len()-1)
	h.data = h.data[:h.Len()-1]
	h.DownAdjust(0)
	return x
}

func nthUglyNumber(n int) int {
	isVisit := make(map[int]bool)
	h := &minHeap{data: make([]int, 0)}
	h.Push(1)
	isVisit[1] = true
	multiples := []int{2, 3, 5}
	var result int
	for n > 0 {
		result = h.Pop()
		n--
		for _, mul := range multiples {
			if !isVisit[result*mul] {
				isVisit[result*mul] = true
				h.Push(result * mul)
			}
		}
	}
	return result
}
