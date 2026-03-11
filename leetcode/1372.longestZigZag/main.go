package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var maxLen int
	var dfs func(*TreeNode) (int, int)
	dfs = func(root *TreeNode) (int, int) { // 分别返回以当前节点为根节点，往左走和往右走的最大交错路径长度
		if root == nil {
			return -1, -1
		}
		_, leftRight := dfs(root.Left)
		rightLeft, _ := dfs(root.Right)
		curLeft := 1 + leftRight
		curRight := 1 + rightLeft
		if curLeft > maxLen {
			maxLen = curLeft
		}
		if curRight > maxLen {
			maxLen = curRight
		}
		return curLeft, curRight
	}
	dfs(root)
	return maxLen
}
