package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum1(root *TreeNode, targetSum int) int {
	// 统计以当前节点为根节点的路径和为targetSum的路径数
	if root == nil {
		return 0
	}
	return returnPathSum(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
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
	return returnPathSum(root.Left, targetSum-root.Val) + returnPathSum(root.Right, targetSum-root.Val) + count
}

func pathSum(root *TreeNode, targetSum int) int {
	// 统计以当前节点为根节点的路径和为targetSum的路径数
	preSum := map[int]int{0: 1} // 前缀和，初始化时有一个0的前缀和
	var count int
	var preOrder func(*TreeNode, int)
	preOrder = func(node *TreeNode, curSum int) {
		if node == nil {
			return
		}
		curSum += node.Val
		count += preSum[curSum-targetSum] // 统计当前节点的前缀和
		preSum[curSum]++                  // 更新前缀和
		preOrder(node.Left, curSum)
		preOrder(node.Right, curSum)
		preSum[curSum]-- // 回溯，恢复前缀和
		return
	}
	preOrder(root, 0)
	return count
}
