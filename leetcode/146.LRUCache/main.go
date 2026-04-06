package main

import "container/list"

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
	}
	return -1
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
