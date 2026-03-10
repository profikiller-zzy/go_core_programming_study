package main

import "fmt"

// ListNode :Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList1(head *ListNode) *ListNode {
	return mergeSort(head, nil)
}

// mergeSort 归并排序，对从head开始到tail结尾，但是不包括tail的链表进行排序
func mergeSort(head, tail *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == tail { // 这里不包含tail，所以head.Next == tail，说明只有一个节点，需要断开再递归排序
		head.Next = nil
		return head
	}

	// 快慢指针找到中点
	slow, fast := head, head
	for fast != tail && fast.Next != tail {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 2. 递归排序左右两部分
	left := mergeSort(head, slow)
	right := mergeSort(slow, tail)

	// 3. 合并
	return merge(left, right)
}

// merge 合并两个升序链表
func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	// 简化剩余部分连接
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}
	return dummy.Next
}

// sortList 迭代方案
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var length int
	node := head
	for node != nil {
		length++
		node = node.Next
	}
	dummyHead := &ListNode{Next: head}

	for step := 1; step < length; step <<= 1 {
		prev := dummyHead
		cur := dummyHead.Next
		for cur != nil {
			head1 := cur
			for i := 1; i < step && cur.Next != nil; i++ {
				cur = cur.Next
			}
			head2 := cur.Next
			cur.Next = nil

			cur = head2
			for i := 1; i < step && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}
			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}

			merged := merge(head1, head2)
			prev.Next = merged
			for prev.Next != nil {
				prev = prev.Next
			}

			cur = next
		}
	}
	return dummyHead.Next
}

// merge 函数同上，略

func main() {
	head := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	newHead := sortList(head)
	for newHead != nil {
		fmt.Println(newHead.Val)
		newHead = newHead.Next
	}
}
