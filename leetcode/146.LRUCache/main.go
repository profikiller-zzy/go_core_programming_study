package main

// https://leetcode.cn/problems/lru-cache/description

type LRUCache struct {
	keyValue map[int]int
	keyOrder []int
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		keyValue: make(map[int]int),
		keyOrder: make([]int, 0, capacity),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if value, ok := this.keyValue[key]; ok {
		// 更新key的使用顺序
		for index, curKey := range this.keyOrder {
			if curKey == key {
				// 将当前key移动到最后
				this.keyOrder = append(this.keyOrder[:index], this.keyOrder[index+1:]...)
				this.keyOrder = append(this.keyOrder, key)
			}
		}
		return value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.keyValue[key]; ok { // 如果key已经存在，更新值
		this.keyValue[key] = value
		// 更新key的使用顺序
		for index, curKey := range this.keyOrder {
			if curKey == key {
				// 将当前key移动到最后
				this.keyOrder = append(this.keyOrder[:index], this.keyOrder[index+1:]...)
				this.keyOrder = append(this.keyOrder, key)
			}
		}
	} else { // key 不存在
		if len(this.keyValue) >= this.capacity {
			delete(this.keyValue, this.keyOrder[0])
			this.keyOrder = this.keyOrder[1:]
		}
		this.keyValue[key] = value
		this.keyOrder = append(this.keyOrder, key)
	}
}
