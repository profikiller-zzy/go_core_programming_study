package main

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 增加一个ghost节点用于指向头节点
	ghostNode := &ListNode{Next: head}
	left := ghostNode
	mid := head
	right := head.Next
	for mid != nil && right != nil {
		left.Next = right
		mid.Next = right.Next
		right.Next = mid

		left = left.Next.Next
		mid = mid.Next
		if mid != nil {
			right = mid.Next
		}
	}
	return ghostNode.Next
}

// 创建链表函数
func createLinkedList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	// 初始化头节点
	head := &ListNode{Val: values[0]}
	current := head

	// 依次添加节点
	for _, value := range values[1:] {
		current.Next = &ListNode{Val: value}
		current = current.Next
	}

	return head
}

// 打印链表函数
func printLinkedList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Println()
}

func main() {
	var tarNode *ListNode = createLinkedList([]int{1, 2, 3, 4})
	printLinkedList(tarNode)
	newHead := swapPairs(tarNode)
	printLinkedList(newHead)
}
