package list

type DoubleListNode struct {
	key  int
	pre  *DoubleListNode
	next *DoubleListNode
}

type LRUCache struct {
	keyValue  map[int]int
	dummyHead *DoubleListNode
	dummyTail *DoubleListNode
	capacity  int
	size      int
}

func Constructor(capacity int) LRUCache {
	head := &DoubleListNode{}
	tail := &DoubleListNode{}
	head.pre = nil
	head.next = tail
	tail.pre = head
	tail.next = nil
	return LRUCache{
		keyValue:  make(map[int]int),
		dummyHead: head,
		dummyTail: tail,
		capacity:  capacity,
		size:      0,
	}
}

func (this *LRUCache) Get(key int) int {
	if value, ok := this.keyValue[key]; ok {
		for cur := this.dummyHead.next; cur != this.dummyTail; cur = cur.next {
			if cur.key == key {
				// 将当前key移动到最后
				cur.pre.next = cur.next
				cur.next.pre = cur.pre
				this.dummyTail.pre.next = cur
				cur.pre = this.dummyTail.pre
				this.dummyTail.pre = cur
				cur.next = this.dummyTail
				break
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
		for cur := this.dummyHead.next; cur != this.dummyTail; cur = cur.next {
			if cur.key == key {
				// 将当前key移动到最后
				cur.pre.next = cur.next
				cur.next.pre = cur.pre
				this.dummyTail.pre.next = cur
				cur.pre = this.dummyTail.pre
				this.dummyTail.pre = cur
				cur.next = this.dummyTail
				break
			}
		}
	} else { // key 不存在
		if this.size >= this.capacity { // 如果缓存已满，删除最久未使用的节点
			delete(this.keyValue, this.dummyHead.next.key)
			this.dummyHead.next = this.dummyHead.next.next
			this.dummyHead.next.pre = this.dummyHead
			this.size--
		}
		// 添加新节点到尾部
		newNode := &DoubleListNode{
			key:  key,
			pre:  nil,
			next: nil,
		}
		newNode.pre = this.dummyTail.pre
		newNode.next = this.dummyTail
		this.dummyTail.pre.next = newNode
		this.dummyTail.pre = newNode
		this.keyValue[key] = value
		this.size++
	}
}
