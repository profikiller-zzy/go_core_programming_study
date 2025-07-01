package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{
		Next: head,
	}
	var tail *ListNode
	var listLen int
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Next == nil {
			tail = cur
		}
		listLen++
	}
	k = k % listLen
	if k == 0 {
		return dummy.Next
	}

	left, right := dummy, dummy
	for i := 0; i < k; i++ {
		right = right.Next
	}
	for right.Next != nil {
		left = left.Next
		right = right.Next
	}
	dummy.Next = left.Next
	left.Next = nil
	tail.Next = head
	return dummy.Next
}
