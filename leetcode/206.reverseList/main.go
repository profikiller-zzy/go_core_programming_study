package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/reverse-linked-list/description

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	cur := head.Next
	head.Next = nil
	for cur != nil {
		next := cur.Next
		cur.Next = head
		head = cur
		cur = next
	}
	return head
}
