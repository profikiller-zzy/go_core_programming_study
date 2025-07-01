package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}
	prev := dummy
	cur := head
	for cur != nil {
		if cur.Next == nil || cur.Val != cur.Next.Val {
			prev = cur
			cur = cur.Next
			continue
		}
		curVal := cur.Val
		for cur != nil && cur.Val == curVal {
			cur = cur.Next
		}
		prev.Next = cur
	}
	return dummy.Next
}
