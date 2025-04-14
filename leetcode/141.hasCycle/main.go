package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/linked-list-cycle/description/
// hasCycle1 哈希表解法
func hasCycle1(head *ListNode) bool {
	nodeHash := make(map[*ListNode]bool)
	for cur := head; cur != nil; cur = cur.Next {
		if _, ok := nodeHash[cur]; ok {
			return true
		} else {
			nodeHash[cur] = true
		}
	}
	return false
}

// hasCycle 使用快慢指针
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func main() {

}
