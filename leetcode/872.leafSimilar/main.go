package main

import (
	"fmt"
	"reflect"
)

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	postOrder := func(root *TreeNode) []int {
		result := make([]int, 0)
		if root == nil {
			return result
		}
		var prev *TreeNode
		cur := root
		stack := make([]*TreeNode, 0)
		for cur != nil || len(stack) > 0 {
			for cur != nil {
				stack = append(stack, cur)
				cur = cur.Left
			}
			top := stack[len(stack)-1]
			if top.Right == prev || top.Right == nil {
				if top.Left == nil && top.Right == nil {
					result = append(result, top.Val)
				}
				stack = stack[:len(stack)-1]
				prev = top
			} else {
				cur = top.Right
			}
		}
		return result
	}
	leafVals1, leafVals2 := postOrder(root1), postOrder(root2)
	return reflect.DeepEqual(leafVals1, leafVals2)
}

func main() {
	root1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	root2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}}
	fmt.Println(leafSimilar(root1, root2))
}
