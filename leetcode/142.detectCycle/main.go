package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/linked-list-cycle-ii/description/
// detectCycle 哈希表
func detectCycle(head *ListNode) *ListNode {
	nodeHash := make(map[*ListNode]bool)
	for cur := head; cur != nil; cur = cur.Next {
		if _, ok := nodeHash[cur]; ok {
			return cur
		} else {
			nodeHash[cur] = true
		}
	}
	return nil
}
