package main

import "fmt"

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	var curValue int
	var preorder func(r *TreeNode)
	preorder = func(r *TreeNode) {
		if r.Right != nil {
			preorder(r.Right)
		}
		r.Val += curValue
		curValue = r.Val
		if r.Left != nil {
			preorder(r.Left)
		}
	}
	preorder(root)
	return root
}

func main() {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	bstToGst(root)
	fmt.Println(root)
}
