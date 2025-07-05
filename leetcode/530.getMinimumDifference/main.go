package main

import "math"

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}
	minDis := math.MaxInt32
	stack := make([]*TreeNode, 0)
	cur, prev := root, (*TreeNode)(nil)
	for cur != nil || len(stack) > 0 {
		for node := cur; node != nil; node = node.Left {
			stack = append(stack, node)
		}
		curVal := stack[len(stack)-1].Val
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil {
			dis := curVal - prev.Val
			if dis < minDis {
				minDis = dis
			}
		}
		prev = cur
		cur = cur.Right
	}
	return minDis
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
