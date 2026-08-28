package main

import "math"

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST1(root *TreeNode) bool {
	stack := make([]*TreeNode, 0)
	prev := math.MinInt64
	cur := root
	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.Val <= prev {
			return false
		}
		prev = cur.Val
		cur = cur.Right
	}
	return true
}

// isValidBST 递归法，递归中序遍历得到遍历数组，最后检查数组是否递增
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var value = make([]int, 0)
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		value = append(value, node.Val)
		traverse(node.Right)
	}
	traverse(root)

	for index := 1; index < len(value); index++ {
		if value[index] <= value[index-1] {
			return false
		}
	}
	return true
}
