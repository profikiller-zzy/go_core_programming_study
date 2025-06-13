package main

type RandomizedSet struct {
	hash  map[int]int
	value []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		hash:  make(map[int]int),
		value: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.hash[val]; ok {
		return false
	}
	this.hash[val] = len(this.value)
	this.value = append(this.value, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.hash[val]
	if !ok { // 元素不存在
		return false
	}
	tailIndex := len(this.value) - 1
	this.value[index] = this.value[tailIndex]
	this.hash[this.value[index]] = index
	this.value = this.value[:tailIndex]
	delete(this.hash, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.value[rand.Intn(len(this.value))]
}
