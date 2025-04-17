package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 使用先序遍历访问树节点
	if root == nil {
		return nil
	}
	if root == p || root == q {
		// 先序遍历首先访问根结点
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	// 如果左右两个返回的节点都不为空，说明p和q分别在左右子树中，而这个root就是p和q的共同祖先
	return root
}
