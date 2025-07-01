package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	if len(inorder) == 1 {
		return &TreeNode{
			Val: inorder[0],
		}
	}
	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{
		Val: rootVal,
	}

	var rootIndex int
	for rootIndex = 0; rootIndex < len(inorder); rootIndex++ {
		if inorder[rootIndex] == rootVal {
			break
		}
	}
	root.Left = buildTree(inorder[:rootIndex], postorder[:rootIndex])
	root.Right = buildTree(inorder[rootIndex+1:], postorder[rootIndex:len(postorder)-1])
	return root
}
