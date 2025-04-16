package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

}

func isAncestor(root, p, q *TreeNode) bool {
	if root == nil {
		return false
	}
	if root == p || root == q {
		return true
	}
	return isAncestor(root.Left, p, q) || isAncestor(root.Right, p, q)
}
