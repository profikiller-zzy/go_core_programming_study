package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	// 其实就是层序遍历
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	curLevel := make([]*TreeNode, 0)
	curLevel = append(curLevel, root)
	for {
		res = append(res, curLevel[len(curLevel)-1].Val)
		nextLevel := make([]*TreeNode, 0)
		for _, node := range curLevel {
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		if len(nextLevel) == 0 {
			return res
		} else {
			curLevel = nextLevel
		}
	}
}
