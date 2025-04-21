package main

import "fmt"

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	// 以当前节点为根节点的最大路径和，也就是路径经过当前节点，并且当前节点的为该路径的根
	// 那么这个最大的和就是当前节点的值加上左右子树的最大路径的和
	if root == nil {
		return 0
	}
	maxVal := root.Val
	var downward func(*TreeNode)
	downward = func(node *TreeNode) {
		if node == nil {
			return
		}
		var curVal int
		curVal = node.Val + returnMaxPathSum(node.Left) + returnMaxPathSum(node.Right)
		if curVal > maxVal {
			maxVal = curVal
		}
		downward(node.Left)
		downward(node.Right)
	}
	downward(root)
	return maxVal
}

func returnMaxPathSum(root *TreeNode) int {
	// 返回以当前节点为根节点的路径和
	var maxVal int
	var downward func(*TreeNode, int)
	downward = func(node *TreeNode, curVal int) {
		if node == nil {
			return
		}
		curVal += node.Val
		if curVal > maxVal {
			maxVal = curVal
		}
		downward(node.Left, curVal)
		downward(node.Right, curVal)
	}
	downward(root, 0)
	return maxVal
}

func main() {
	//root := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val: 2,
	//	},
	//	Right: &TreeNode{
	//		Val: 3,
	//	},
	//}
	root := &TreeNode{
		Val: -3,
	}
	fmt.Println(maxPathSum(root))
}
