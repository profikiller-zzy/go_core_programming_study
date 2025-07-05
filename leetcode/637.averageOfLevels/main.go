package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	result := make([]float64, 0)
	curLevel := []*TreeNode{root}
	for len(curLevel) > 0 {
		var curLevelSum float64
		nextLevel := make([]*TreeNode, 0)
		for _, node := range curLevel {
			curLevelSum += float64(node.Val)
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		result = append(result, curLevelSum/float64(len(curLevel)))
		curLevel = nextLevel
	}
	return result
}
