package main

import "math"

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	return dfs(root, math.MinInt32)
}

func dfs(root *TreeNode, maxVal int) int {
	if root == nil {
		return 0
	}
	var res int
	if root.Val >= maxVal {
		maxVal = root.Val
		res++
	}
	res += dfs(root.Left, maxVal) + dfs(root.Right, maxVal)
	return res
}
