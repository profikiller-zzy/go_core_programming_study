package main

import "math"

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	level := 1
	currentLevel := []*TreeNode{root}
	maxSum := math.MinInt32
	maxLevel := 1
	for len(currentLevel) > 0 {
		nextLevel := make([]*TreeNode, 0)
		var curSum int
		for _, node := range currentLevel {
			curSum += node.Val
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		if curSum > maxSum {
			maxSum = curSum
			maxLevel = level
		}
		currentLevel = nextLevel
		level++
	}
	return maxLevel
}
