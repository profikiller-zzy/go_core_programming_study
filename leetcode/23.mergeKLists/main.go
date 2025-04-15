package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/merge-k-sorted-lists/description

func mergeKLists(lists []*ListNode) *ListNode {
	var res *ListNode
	if len(lists) == 0 {
		return res
	} else if len(lists) == 1 {
		return lists[0]
	}

	// 归并两个链表，返回归并后的头节点
	var merge2Lists func(*ListNode, *ListNode) *ListNode
	merge2Lists = func(node1 *ListNode, node2 *ListNode) *ListNode {
		if node1 == nil {
			return node2
		} else if node2 == nil {
			return node1
		}
		var resHead, resTail *ListNode

		// 初始化结果链表的头节点
		if node1.Val <= node2.Val {
			resHead = node1
			node1 = node1.Next
		} else {
			resHead = node2
			node2 = node2.Next
		}
		resTail = resHead

		for node1 != nil && node2 != nil {
			if node1.Val <= node2.Val {
				resTail.Next = node1
				node1 = node1.Next
			} else {
				resTail.Next = node2
				node2 = node2.Next
			}
			resTail = resTail.Next
		}

		// 将剩余的列表处理了
		if node1 == nil {
			resTail.Next = node2
		} else {
			resTail.Next = node1
		}
		return resHead
	}

	// 不断两两合并直到最后合并为一个长链表
	for len(lists) > 1 {
		var newLists []*ListNode
		for i := 0; i < len(lists); i += 2 {
			if i+1 < len(lists) { // 假如这一轮有两个列表
				merged := merge2Lists(lists[i], lists[i+1])
				newLists = append(newLists, merged)
			} else { // 如果只有一个链表，直接添加不用合并
				newLists = append(newLists, lists[i])
			}
		}
		lists = newLists
	}

	return lists[0]
}

func main() {

}
