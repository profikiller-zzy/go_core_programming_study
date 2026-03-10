package main

import (
	"fmt"
	"math"
)

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32              // 初始化为最小整数
	var maxGain func(node *TreeNode) int // 计算以当前节点为根节点的单向最大路径和

	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 递归计算左右子树单向最大路径和（负值则忽略）
		leftGain := max(0, maxGain(node.Left))
		rightGain := max(0, maxGain(node.Right))

		// 以当前节点为根的完整路径和（可能为最大值）
		currentPath := node.Val + leftGain + rightGain
		if currentPath > maxSum {
			maxSum = currentPath
		}

		// 返回单向最大路径和（用于父节点计算）
		return node.Val + max(leftGain, rightGain)
	}

	maxGain(root)
	return maxSum
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
