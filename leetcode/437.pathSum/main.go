package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {

}

func returnPathSum(root *TreeNode, targetSum int) int {
	// 返回以当前节点为根节点的路径和为targetSum的路径数
	if root == nil {
		return 0
	}
	var count int
	if root.Val == targetSum {
		count++
	}
	return returnPathSum(root.Left, targetSum-root.Val) + returnPathSum(root.Right, targetSum-root.Val)
}
