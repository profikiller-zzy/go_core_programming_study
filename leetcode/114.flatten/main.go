package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	_ = flattenTree(root)
	return
}

// flattenTree 递归实现
// 时间复杂度O(n)，空间复杂度O(H)，H为树的高度
func flattenTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := flattenTree(root.Left)
	right := flattenTree(root.Right)

	cur := root
	cur.Left = nil
	if left != nil {
		cur.Right = left
		for cur.Right != nil {
			cur = cur.Right
		}
	}
	cur.Right = right
	return root
}

// flatten1 使用寻找前驱节点的方法
// 时间复杂度O(n)，空间复杂度O(1)
func flatten1(root *TreeNode) {
	if root == nil {
		return
	}
	cur := root
	for cur != nil {
		if cur.Left != nil {
			next := cur.Left
			pre := cur.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = cur.Right
			cur.Left = nil
			cur.Right = next
		}
		cur = cur.Right
	}
}
