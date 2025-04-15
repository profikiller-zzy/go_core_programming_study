package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0

	var depth func(*TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftDepth := depth(node.Left)
		rightDepth := depth(node.Right)

		// 更新最大直径
		maxDiameter = max(maxDiameter, leftDepth+rightDepth)

		// 当前节点的深度
		return max(leftDepth, rightDepth) + 1
	}

	depth(root)
	return maxDiameter
}
