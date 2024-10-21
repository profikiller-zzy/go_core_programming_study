package main

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head} // 创建一个哑节点，指向head
	left, right := dummy, dummy    // 将left和right初始化为哑节点

	// 让right先走n步
	for i := 0; i < n; i++ {
		right = right.Next
	}

	// 使用双指针法找到倒数第n+1个指针
	for right.Next != nil {
		left = left.Next
		right = right.Next
	}

	// 删除第n个节点(GO不需要显式释放内存)
	left.Next = left.Next.Next

	// 返回head，注意head可能已被删除，所以返回dummy.Next
	return dummy.Next
}

func main() {

}
