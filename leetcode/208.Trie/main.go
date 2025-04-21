package main

import "fmt"

type Trie struct {
	byte     byte           // 当前节点存储的字符
	children map[byte]*Trie // 子节点
	isEnd    bool
}

func Constructor() Trie {
	// 返回根节点
	return Trie{
		byte: byte('0'),
	}
}

func (this *Trie) Insert(word string) {
	cur := this
	for index := 0; index < len(word); index++ {
		ch := word[index]
		if cur.children == nil { // 如果没有子节点，则初始化
			cur.children = make(map[byte]*Trie)
		}
		var child *Trie
		if _, ok := cur.children[ch]; !ok { // 没有则添加
			child = &Trie{byte: ch}
			cur.children[ch] = child
		} else { // 有的话直接取
			child = cur.children[ch]
		}
		cur = cur.children[ch]
	}
	cur.isEnd = true
}

func (this *Trie) Search(word string) bool {
	cur := this
	for index := 0; index < len(word); index++ {
		ch := word[index]
		if cur.children == nil {
			return false
		}

		if node, ok := cur.children[ch]; ok {
			cur = node
		} else { // 没有返回false
			return false
		}
	}
	if cur.isEnd {
		return true
	} else {
		return false
	}
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for index := 0; index < len(prefix); index++ {
		ch := prefix[index]
		if cur.children == nil {
			return false
		}

		if node, ok := cur.children[ch]; ok {
			cur = node
		} else { // 没有返回false
			return false
		}
	}
	return true
}

func main() {
	trie := Constructor()
	root := &trie
	root.Insert("apple")
	fmt.Println(root.Search("apple"))   // 返回 True
	fmt.Println(root.Search("app"))     // 返回 False
	fmt.Println(root.StartsWith("app")) // 返回 True
	root.Insert("app")
	fmt.Println(root.Search("app")) // 返回 True
}
