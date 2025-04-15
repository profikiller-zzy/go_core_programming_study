package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)
	currentLevel := []*TreeNode{root} // 用于存储当前这层的节点
	for {
		nextLevel := make([]*TreeNode, 0)
		currentLevelVal := make([]int, 0)
		for _, node := range currentLevel {
			currentLevelVal = append(currentLevelVal, node.Val)
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		result = append(result, currentLevelVal)
		if len(nextLevel) == 0 {
			return result
		} else {
			currentLevel = nextLevel
		}
	}
}
