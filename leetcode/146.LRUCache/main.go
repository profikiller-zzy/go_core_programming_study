package main

import "container/list"

//type LRUCache struct {
//	keyValue map[int]int
//	keyOrder []int
//	capacity int
//}
//
//func Constructor(capacity int) LRUCache {
//	return LRUCache{
//		keyValue: make(map[int]int),
//		keyOrder: make([]int, 0, capacity),
//		capacity: capacity,
//	}
//}
//
//func (this *LRUCache) Get(key int) int {
//	if value, ok := this.keyValue[key]; ok {
//		// 更新key的使用顺序
//		for index, curKey := range this.keyOrder {
//			if curKey == key {
//				// 将当前key移动到最后
//				this.keyOrder = append(this.keyOrder[:index], this.keyOrder[index+1:]...)
//				this.keyOrder = append(this.keyOrder, key)
//			}
//		}
//		return value
//	} else {
//		return -1
//	}
//}
//
//func (this *LRUCache) Put(key int, value int) {
//	if _, ok := this.keyValue[key]; ok { // 如果key已经存在，更新值
//		this.keyValue[key] = value
//		// 更新key的使用顺序
//		for index, curKey := range this.keyOrder {
//			if curKey == key {
//				// 将当前key移动到最后
//				this.keyOrder = append(this.keyOrder[:index], this.keyOrder[index+1:]...)
//				this.keyOrder = append(this.keyOrder, key)
//			}
//		}
//	} else { // key 不存在
//		if len(this.keyValue) >= this.capacity {
//			delete(this.keyValue, this.keyOrder[0])
//			this.keyOrder = this.keyOrder[1:]
//		}
//		this.keyValue[key] = value
//		this.keyOrder = append(this.keyOrder, key)
//	}
//}

type LRUCache struct {
	capacity int
	list     *list.List            // 双向链表
	kvMap    map[int]*list.Element // key -> *list.Element 的 map
}

type entry struct {
	key, value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity, list.New(), map[int]*list.Element{}}
}

func (c *LRUCache) Get(key int) int {
	if elem, ok := c.kvMap[key]; ok { // 有这本书
		c.list.MoveToFront(elem) // 把这本书放在最上面
		return elem.Value.(*entry).value
	} else {
		return -1
	}
}

func (c *LRUCache) Put(key, value int) {
	if elem, ok := c.kvMap[key]; ok { // 有这本书
		elem.Value.(*entry).value = value // 更新
		c.list.MoveToFront(elem)          // 把这本书放在最上面
		return
	}
	c.kvMap[key] = c.list.PushFront(&entry{key, value}) // 新书，放在最上面
	if len(c.kvMap) > c.capacity {                      // 书太多了
		delete(c.kvMap, c.list.Remove(c.list.Back()).(*entry).key) // 去掉最后一本书
	}
}
