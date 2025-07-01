package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	curLevel := []*TreeNode{root}
	sums := make([]int, 0)
	for len(curLevel) > 0 {
		nextLevel := make([]*TreeNode, 0)
		for _, node := range curLevel {
			if node.Left == nil && node.Right == nil {
				sums = append(sums, node.Val)
			}
			if node.Left != nil {
				node.Left.Val += node.Val * 10
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				node.Right.Val += node.Val * 10
				nextLevel = append(nextLevel, node.Right)
			}
		}
		curLevel = nextLevel
	}
	total := 0
	for _, sum := range sums {
		total += sum
	}
	return total
}
