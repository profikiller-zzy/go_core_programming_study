package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	dummy2 := &ListNode{}
	prev, cur := dummy, head
	p2tail := dummy2
	for cur != nil {
		if cur.Val < x {
			prev = prev.Next
			cur = cur.Next
			continue
		}
		p2tail.Next = cur
		prev.Next = cur.Next
		p2tail = cur
		cur.Next = nil
		cur = prev.Next
	}
	prev.Next = dummy2.Next
	return dummy.Next
}
