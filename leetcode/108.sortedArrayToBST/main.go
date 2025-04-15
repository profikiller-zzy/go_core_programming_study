package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// sortedArrayToBST 利用了bst的中序遍历是升序序列的特点，相当于是将一个升序序列还原成bst
// 每次选择中间节点作为根节点，左边的序列作为左子树，右边的序列作为右子树，不断递归
func sortedArrayToBST(nums []int) *TreeNode {
	return dfs(nums, 0, len(nums)-1)
}

func dfs(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	// 以升序序列的中间当作根节点
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = dfs(nums, left, mid-1)
	root.Right = dfs(nums, mid+1, right)
	return root
}
