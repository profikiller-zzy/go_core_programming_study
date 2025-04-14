package main

// Node :Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// https://leetcode.cn/problems/copy-list-with-random-pointer/description

// copyRandomList1 使用到了map存储原本的节点和新节点的映射关系
func copyRandomList1(head *Node) *Node {
	if head == nil {
		return nil
	}
	hashNode := make(map[*Node]*Node) // 使用map存储原本的节点和新节点的映射关系
	newDummy := &Node{Val: -1, Next: nil}
	newCur := newDummy

	// 第一遍遍历，复制节点
	for cur := head; cur != nil; cur = cur.Next {
		newNext := &Node{Val: cur.Val}
		hashNode[cur] = newNext
		newCur.Next = newNext
		newCur = newNext
	}

	// 第二次遍历，复制随机指针
	newCur = newDummy.Next
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Random != nil {
			newCur.Random = hashNode[cur.Random]
		} else {
			newCur.Random = nil
		}
		newCur = newCur.Next
	}
	return newDummy.Next
}

// copyRandomList 迭代加节点拆分
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	for cur := head; cur != nil; cur = cur.Next.Next {
		cur.Next = &Node{Val: cur.Val, Next: cur.Next}
	}
	for cur := head; cur != nil; cur = cur.Next.Next {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
	}

	// 将深拷贝的链表和原链表分开
	newHead := head.Next
	for cur := head; cur != nil; cur = cur.Next {
		next := cur.Next
		cur.Next = cur.Next.Next
		if next.Next != nil {
			next.Next = next.Next.Next
		}
	}
	return newHead
}
