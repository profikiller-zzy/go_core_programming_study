package mergeTwoLists

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/merge-two-sorted-lists/description

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	cur1, cur2 := list1, list2
	var (
		newHead *ListNode
		newCur  *ListNode
	)
	if cur1.Val <= cur2.Val {
		newHead = cur1
		cur1 = cur1.Next
	} else {
		newHead = cur2
		cur2 = cur2.Next
	}
	newCur = newHead
	for cur1 != nil && cur2 != nil {
		if cur1.Val <= cur2.Val {
			newCur.Next = cur1
			cur1 = cur1.Next
		} else {
			newCur.Next = cur2
			cur2 = cur2.Next
		}
		newCur = newCur.Next
	}
	if cur1 == nil {
		newCur.Next = cur2
	} else {
		newCur.Next = cur1
	}
	return newHead
}

func main() {

}
