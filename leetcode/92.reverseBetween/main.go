package main

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	leftNode, rightNode := dummy, dummy
	var (
		preLeft   *ListNode
		nextRight *ListNode
	)
	for left != 0 {
		if left == 1 {
			preLeft = leftNode
		}
		leftNode = leftNode.Next
		left--
	}
	for right != 0 {
		rightNode = rightNode.Next
		right--
	}
	nextRight = rightNode.Next

	reverse(leftNode, rightNode)
	preLeft.Next = rightNode
	leftNode.Next = nextRight
	return dummy.Next
}

func reverse(head *ListNode, tail *ListNode) {
	var prev *ListNode
	cur := head
	for cur != tail {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	tail.Next = prev
}

func main() {
	// Example usage:
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	newHead := reverseBetween(head, 2, 4)
	fmt.Println(newHead) // Output the modified list
}
