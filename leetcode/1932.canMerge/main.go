package main

import "fmt"

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func canMerge(trees []*TreeNode) *TreeNode {
	if len(trees) == 1 {
		return trees[0]
	}
	rootMap := make(map[int]*TreeNode)
	for _, tree := range trees {
		rootMap[tree.Val] = tree
	}

	leafMap := make(map[int]*TreeNode)
	for _, tree := range trees {
		findLeaves(tree, nil, false, leafMap)
	}
	for len(trees) > 1 {
		merged := false
		for _, tree := range trees {
			if parent, exists := leafMap[tree.Val]; exists && parent != nil {
				if parent.Left != nil && parent.Left.Val == tree.Val {
					parent.Left = tree
				} else if parent.Right != nil && parent.Right.Val == tree.Val {
					parent.Right = tree
				}
				for i, t := range trees {
					if t.Val == tree.Val {
						trees = append(trees[:i], trees[i+1:]...)
						break
					}
				}

				delete(rootMap, tree.Val)
				merged = true
				break
			}
		}

		if !merged {
			return nil
		}
	}
	if len(trees) == 1 && isValidBST(trees[0]) {
		return trees[0]
	}

	return nil
}

// 查找所有叶节点并记录其父节点
func findLeaves(root, parent *TreeNode, isLeft bool, leafMap map[int]*TreeNode) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		leafMap[root.Val] = parent
		return
	}
	findLeaves(root.Left, root, true, leafMap)
	findLeaves(root.Right, root, false, leafMap)
}

func isValidBST(root *TreeNode) bool {
	var prev *TreeNode
	var stack []*TreeNode
	cur := root
	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil && prev.Val >= cur.Val {
			return false
		}
		prev = cur
		cur = cur.Right
	}
	return true
}

func main() {
	t1 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
	}
	t2 := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 5},
	}
	t3 := &TreeNode{
		Val:  5,
		Left: &TreeNode{Val: 4},
	}
	fmt.Println(canMerge([]*TreeNode{t1, t2, t3}))
}
