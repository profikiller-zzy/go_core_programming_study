package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	Nodes []*TreeNode
	size  int
}

func newStack() *Stack {
	return &Stack{
		Nodes: make([]*TreeNode, 0),
		size:  0,
	}
}

func (stack *Stack) push(node *TreeNode) {
	stack.Nodes = append(stack.Nodes, node)
	stack.size++
}

func (stack *Stack) pop() *TreeNode {
	if stack.size == 0 {
		return nil
	}
	node := stack.Nodes[stack.size-1]
	stack.Nodes = stack.Nodes[:stack.size-1]
	stack.size--
	return node
}

func kthSmallest(root *TreeNode, k int) int {
	stack := newStack()
	cur := root
	for cur != nil || stack.size > 0 {
		for cur != nil {
			stack.push(cur)
			cur = cur.Left
		}
		cur = stack.pop()
		k--
		if k == 0 {
			return cur.Val
		}
		cur = cur.Right
	}
	return -1
}
