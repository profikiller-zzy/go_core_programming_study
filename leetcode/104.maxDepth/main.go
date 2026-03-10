package main

import "container/list"

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	} else {
		return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
	}
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := list.New()
	depth := 0
	queue.PushFront(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			elem := queue.Front()
			queue.Remove(elem)
			node := elem.Value.(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		depth++
	}
	return depth
}
