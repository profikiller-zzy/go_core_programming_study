package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)
	curtLevel := []*TreeNode{root} // 用于存储当前这层的节点
	for len(curtLevel) > 0 {
		curtLevelVal := make([]int, len(curtLevel))
		nextLevel := make([]*TreeNode, 0)
		for index, node := range curtLevel {
			curtLevelVal[index] = node.Val
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		result = append(result, curtLevelVal)
		curtLevel = nextLevel
	}
	for index := 0; index < len(result); index++ {
		if index%2 == 0 {
			continue
		}
		// 奇数层需要反转
		for left, right := 0, len(result[index])-1; left < right; left, right = left+1, right-1 {
			result[index][left], result[index][right] = result[index][right], result[index][left]
		}
	}
	return result
}
