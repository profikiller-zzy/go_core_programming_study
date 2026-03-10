package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	cur := root
	for cur != nil {
		if cur.Left != nil {
			leftRoot := cur.Left
			leftTail := cur.Left
			for leftTail.Right != nil {
				leftTail = leftTail.Right
			}
			leftTail.Right = cur.Right
			cur.Left = nil
			cur.Right = leftRoot
		}
		cur = cur.Right
	}
}
