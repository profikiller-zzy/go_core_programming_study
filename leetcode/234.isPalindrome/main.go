package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/palindrome-linked-list/description

func isPalindrome(head *ListNode) bool {
	// 先寻找中间节点
	if head.Next == nil {
		return true
	}
	mid, fast := head, head
	for fast != nil && fast.Next != nil {
		mid = mid.Next
		fast = fast.Next.Next
	}

	// 反转后半部分
	var prev *ListNode
	cur := mid
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	for prev != nil {
		if prev.Val != head.Val {
			return false
		}
		head = head.Next
		prev = prev.Next
	}
	return true
}
