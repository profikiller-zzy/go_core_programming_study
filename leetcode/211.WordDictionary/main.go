package main

import "container/list"

type WordDictionary struct {
	ch       byte
	children map[byte]*WordDictionary
	isEnd    bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		ch: byte('0'),
	}
}

func (this *WordDictionary) AddWord(word string) {
	cur := this
	for index := 0; index < len(word); index++ {
		ch := word[index]
		if cur.children == nil {
			cur.children = make(map[byte]*WordDictionary)
		}
		if _, ok := cur.children[ch]; !ok { // 当前节点的子节点中并没有这个 byte
			child := &WordDictionary{ch: ch}
			cur.children[ch] = child
		}
		cur = cur.children[ch]
	}
	cur.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.dfs(word, 0)
}

func (this *WordDictionary) dfs(word string, index int) bool {
	if index == len(word) {
		return this.isEnd
	}
	ch := word[index]
	if ch == '.' {
		if this.children != nil {
			for _, child := range this.children {
				if child.dfs(word, index+1) {
					return true
				}
			}
		} else {
			return false
		}
	} else {
		if this.children != nil {
			if child, ok := this.children[ch]; ok {
				return child.dfs(word, index+1)
			}
		} else {
			return false
		}
	}
	return false
}

// bfs 搜索 效率低下
func (this *WordDictionary) Search1(word string) bool {
	cur := this
	curLevel := list.New()
	curLevel.PushBack(cur)
	for index := 0; index < len(word); index++ {
		ch := word[index]
		nextLevel := list.New()
		for curLevel.Len() > 0 {
			cur = curLevel.Remove(curLevel.Front()).(*WordDictionary)
			if cur.children == nil {
				continue
			}
			if ch == '.' {
				for _, child := range cur.children {
					nextLevel.PushBack(child)
				}
			} else {
				if _, ok := cur.children[ch]; ok {
					nextLevel.PushBack(cur.children[ch])
				}
			}
		}
		curLevel = nextLevel
	}
	for curLevel.Len() > 0 {
		cur = curLevel.Remove(curLevel.Front()).(*WordDictionary)
		if cur.isEnd {
			return true
		}
	}
	return false
}
