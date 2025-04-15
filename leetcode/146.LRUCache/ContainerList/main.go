package main

import "container/list"

type LRUCache struct {
	capacity  int
	list      *list.List // 双向链表
	keyToNode map[int]*list.Element
}

type entry struct {
	key, value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity, list.New(), map[int]*list.Element{}}
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.keyToNode[key]; ok { // 没有这本书
		c.list.MoveToFront(node) // 把这本书放在最上面
		return node.Value.(entry).value
	} else {
		return -1
	}
}

func (c *LRUCache) Put(key, value int) {
	if node, ok := c.keyToNode[key]; ok { // 有这本书
		node.Value = entry{key, value} // 更新
		c.list.MoveToFront(node)       // 把这本书放在最上面
		return
	}
	c.keyToNode[key] = c.list.PushFront(entry{key, value}) // 新书，放在最上面
	if len(c.keyToNode) > c.capacity {                     // 书太多了
		delete(c.keyToNode, c.list.Remove(c.list.Back()).(entry).key) // 去掉最后一本书
	}
}
