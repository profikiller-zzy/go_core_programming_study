package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	curLevel := []*TreeNode{root}
	for len(curLevel) > 0 {
		nextLevel := make([]*TreeNode, 0)
		for _, node := range curLevel {
			if node.Left == nil && node.Right == nil {
				if node.Val == targetSum {
					return true
				}
			}
			if node.Left != nil {
				node.Left.Val += node.Val
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				node.Right.Val += node.Val
				nextLevel = append(nextLevel, node.Right)
			}
		}
		curLevel = nextLevel
	}
	return false
}
