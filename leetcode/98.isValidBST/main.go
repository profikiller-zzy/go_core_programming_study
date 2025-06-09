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

func isValidBST1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var (
		prev int
		cur  int
	)
	prev = -1 << 63 // Initialize to the minimum value of int
	stack := newStack()
	curNode := root
	for curNode != nil || stack.size > 0 {
		for curNode != nil {
			stack.push(curNode)
			curNode = curNode.Left
		}
		curNode = stack.pop()
		cur = curNode.Val
		if cur <= prev {
			return false
		} else {
			prev = cur
		}
		curNode = curNode.Right
	}
	return true
}

// isValidBST 递归法，递归中序遍历得到遍历数组，最后检查数组是否递增
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var value = make([]int, 0)
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		value = append(value, node.Val)
		traverse(node.Right)
	}
	traverse(root)

	for index := 1; index < len(value); index++ {
		if value[index] <= value[index-1] {
			return false
		}
	}
	return true
}
