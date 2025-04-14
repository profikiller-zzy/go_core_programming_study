package main

// ListNode :Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/sort-list/description

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

// sort 归并排序，对从head开始到tail结尾，但是不包括tail的链表进行排序
func sort(head, tail *ListNode) *ListNode {
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
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}

	mid := slow
	return merge(sort(head, mid), sort(mid, tail))
}

// merge 合并两个升序链表
func merge(head1, head2 *ListNode) *ListNode {
	dummyNode := &ListNode{}
	cur, cur1, cur2 := dummyNode, head1, head2
	for cur1 != nil && cur2 != nil {
		if cur1.Val <= cur2.Val {
			cur.Next = cur1
			cur1 = cur1.Next
		} else {
			cur.Next = cur2
			cur2 = cur2.Next
		}
		cur = cur.Next
	}
	if cur1 != nil {
		cur.Next = cur1
	} else {
		cur.Next = cur2
	}
	return dummyNode.Next
}

func main() {

}
