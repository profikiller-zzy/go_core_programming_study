package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/intersection-of-two-linked-lists/description

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	hashNode := make(map[*ListNode]bool)
	for curA := headA; curA != nil; curA = curA.Next {
		hashNode[curA] = true
	}
	for curB := headB; curB != nil; curB = curB.Next {
		if _, ok := hashNode[curB]; ok {
			return curB
		}
	}
	return nil
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	for curA != curB {
		if curA != nil {
			curA = curA.Next
		} else {
			curA = headB
		}

		if curB != nil {
			curB = curB.Next
		} else {
			curB = headA
		}
	}
	return curA
}
