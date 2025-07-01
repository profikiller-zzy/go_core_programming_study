package main

import (
	"fmt"
)

// ListNode :Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}
	prev := dummy
	curGroupHead := head

	// 先检查剩余需要处理的节点的个数是否大于group长度
	for {
		var nextGroupHead *ListNode = curGroupHead
		for i := 0; i < k; i++ {
			if nextGroupHead == nil { // 长度不够,直接返回
				return dummy.Next
			}
			nextGroupHead = nextGroupHead.Next
		}
		groupHead, groupTail := reverseGroup(curGroupHead, k)
		prev.Next = groupHead
		groupTail.Next = nextGroupHead
		prev = groupTail
		curGroupHead = nextGroupHead
	}
}

// reverseGroup 辅助函数，将以head为头的链表的前k个节点反转
func reverseGroup(head *ListNode, k int) (*ListNode, *ListNode) {
	var (
		cur       *ListNode = head
		prev      *ListNode = nil
		groupHead *ListNode = nil
		groupTail *ListNode = head
	)
	for i := 0; i < k; i++ {
		if i == k-1 {
			groupHead = cur
		}
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return groupHead, groupTail
}

func main() {
	head := &ListNode{
		Val: 10,
	}
	for i := 9; i > 0; i-- {
		cur := &ListNode{
			Val: i,
		}
		cur.Next = head
		head = cur
	}
	for cur := head; cur != nil; cur = cur.Next {
		fmt.Printf("%d ", cur.Val)
	}

	fmt.Println()
	head = reverseKGroup(head, 2)
	for cur := head; cur != nil; cur = cur.Next {
		fmt.Printf("%d ", cur.Val)
	}
}
