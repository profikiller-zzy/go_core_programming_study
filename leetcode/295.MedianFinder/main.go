package main

type MedianFinder struct {
	data []int
}

func Constructor() MedianFinder {
	return MedianFinder{
		data: make([]int, 0),
	}
}

func (this *MedianFinder) AddNum(num int) {
	this.data = append(this.data, num)
	for i := len(this.data) - 1; i > 0; i-- {
		if this.data[i] < this.data[i-1] {
			this.data[i], this.data[i-1] = this.data[i-1], this.data[i]
		} else {
			break
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.data)%2 == 0 {
		return float64(this.data[len(this.data)/2-1]+this.data[len(this.data)/2]) / 2
	}
	return float64(this.data[len(this.data)/2])
}
