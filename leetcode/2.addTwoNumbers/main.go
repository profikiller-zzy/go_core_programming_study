package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/add-two-numbers/description

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	cur1, cur2 := l1, l2
	newHead := &ListNode{Val: 0}
	newCur := newHead
	var (
		curVal int
		carry  int
	)
	for ; cur1 != nil && cur2 != nil; cur1, cur2 = cur1.Next, cur2.Next {
		val := cur1.Val + cur2.Val + carry
		curVal = val % 10 // 当前位的值
		carry = val / 10  // 进位的值
		newCur.Next = &ListNode{Val: curVal}
		newCur = newCur.Next
	}

	var cur *ListNode
	if cur1 != nil {
		cur = cur1
	} else if cur2 != nil {
		cur = cur2
	} else {
		if carry != 0 {
			newCur.Next = &ListNode{Val: carry}
		}
		return newHead.Next
	}
	for cur != nil {
		val := cur.Val + carry
		curVal = val % 10
		carry = val / 10
		newCur.Next = &ListNode{Val: curVal}
		newCur = newCur.Next
		cur = cur.Next
	}
	if carry != 0 {
		newCur.Next = &ListNode{Val: carry}
	}
	return newHead.Next
}

func main() {

}
