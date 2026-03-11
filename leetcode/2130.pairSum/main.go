package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 翻转slow之后的链表
	prev := slow.Next
	cur := prev.Next
	for cur != nil {
		next := cur.Next
		cur.Next = slow.Next
		slow.Next = cur
		prev.Next = next
		cur = next
	}
	mid := slow.Next
	cur = head
	var maxSum int
	for mid != nil {
		curSum := cur.Val + mid.Val
		if curSum > maxSum {
			maxSum = curSum
		}
		cur = cur.Next
		mid = mid.Next
	}
	return maxSum
}
