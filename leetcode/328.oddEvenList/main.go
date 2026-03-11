package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList1(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	if head == nil || head.Next == nil {
		return dummy.Next
	}
	// 从(dummy, leftPtr] 的部分全是原本奇索引指针, 从(leftPtr, rightPtr) 都是偶索引指针，指向目前待处理的指针
	leftPtr, rightPtr, rightPrev, index := dummy, head, dummy, 1
	for rightPtr != nil {
		if index%2 == 0 {
			rightPtr = rightPtr.Next
			rightPrev = rightPrev.Next
		} else {
			if leftPtr.Next == rightPtr {
				rightPtr = rightPtr.Next
				rightPrev = rightPrev.Next
				leftPtr = leftPtr.Next
			} else {
				rightPrev.Next = rightPtr.Next
				rightPtr.Next = leftPtr.Next
				leftPtr.Next = rightPtr
				leftPtr = leftPtr.Next
				rightPtr = rightPrev.Next
			}
		}
		index++
	}
	return dummy.Next
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odd := head       // 奇数链表尾指针
	even := head.Next // 偶数链表头指针（也是当前的偶数节点）
	evenHead := even  // 保存偶数链表头，用于最后拼接

	// 只要偶数节点不为空，且偶数节点的下一个也不为空
	for even != nil && even.Next != nil {
		// 1. 奇数指针向后跳两格 (连接到下一个奇数)
		odd.Next = even.Next
		odd = odd.Next

		// 2. 偶数指针向后跳两格 (连接到下一个偶数)
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
