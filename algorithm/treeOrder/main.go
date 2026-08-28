package main

import "fmt"

// MultiTreeNode 多叉树
type MultiTreeNode struct {
	Val      int
	Children []*MultiTreeNode
}

// TreeNode 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// MultiPostOrderIteration 迭代版本实现多叉树的后序遍历
func MultiPostOrderIteration(root *MultiTreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := []*MultiTreeNode{root}
	var prev *MultiTreeNode
	result := make([]int, 0)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		if len(top.Children) == 0 || (prev != nil && top.Children[len(top.Children)-1] == prev) {
			// 当前节点是叶子节点或者当前节点的所有子节点都已经访问过了
			prev = top
			stack = stack[:len(stack)-1]
			result = append(result, top.Val)
		} else {
			for i := len(top.Children) - 1; i >= 0; i-- { // 从右往左入栈
				stack = append(stack, top.Children[i])
			}
		}
	}
	return result
}

// MultiPostOrderRecursion 递归版本实现多叉树的后序遍历
func MultiPostOrderRecursion(root *MultiTreeNode) []int {
	result := make([]int, 0)
	var dfs func(node *MultiTreeNode)
	dfs = func(node *MultiTreeNode) {
		if node == nil {
			return
		}
		if len(node.Children) > 0 {
			for _, child := range node.Children {
				dfs(child)
			}
		}
		result = append(result, node.Val)
	}
	dfs(root)
	return result
}

func MultiPreOrderIteration(root *MultiTreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := []*MultiTreeNode{root}
	result := make([]int, 0)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		if len(cur.Children) > 0 {
			for i := len(cur.Children) - 1; i >= 0; i-- {
				stack = append(stack, cur.Children[i])
			}
		}
	}
	return result
}

func MultiPreOrderRecursion(root *MultiTreeNode) []int {
	result := make([]int, 0)
	var dfs func(node *MultiTreeNode)
	dfs = func(node *MultiTreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		if len(node.Children) > 0 {
			for _, child := range node.Children {
				dfs(child)
			}
		}
	}
	dfs(root)
	return result
}

// PostOrderIteration 二叉树的后序遍历迭代实现
func PostOrderIteration(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result []int
	var prev *TreeNode
	cur := root
	var stack []*TreeNode
	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		top := stack[len(stack)-1]
		if (prev != nil && top.Right == prev) || top.Right == nil {
			result = append(result, top.Val)
			prev = top
			stack = stack[:len(stack)-1]
		} else {
			cur = top.Right
		}
	}
	return result
}

// PostOrderRecursion 二叉树的后序遍历递归实现
func PostOrderRecursion(root *TreeNode) []int {
	result := make([]int, 0)
	var postOrder func(node *TreeNode)
	postOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postOrder(root.Left)
		postOrder(root.Right)
		result = append(result, node.Val)
	}
	postOrder(root)
	return result
}

// InorderIteration 二叉树的中序遍历迭代实现
func InorderIteration(root *TreeNode) []int {
	result := make([]int, 0)
	cur := root
	var stack []*TreeNode
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, top.Val)
		cur = top.Right
	}
	return result
}

// InorderRecursion 二叉树的中序遍历递归实现
func InorderRecursion(root *TreeNode) []int {
	result := make([]int, 0)
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		result = append(result, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return result
}

func PreorderIteration(root *TreeNode) []int {
	result := make([]int, 0)
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}
	return result
}

func PreorderRecursion(root *TreeNode) []int {
	result := make([]int, 0)
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return result
}

func main() {
	root := &MultiTreeNode{Val: 1}
	root.Children = []*MultiTreeNode{
		{Val: 2},
		{Val: 3},
		{Val: 4},
	}
	root.Children[0].Children = []*MultiTreeNode{
		{Val: 5},
	}
	root.Children[0].Children[0].Children = []*MultiTreeNode{
		{Val: 6},
		{Val: 7},
	}
	fmt.Println(MultiPostOrderIteration(root))
	fmt.Println(MultiPostOrderRecursion(root))
	fmt.Println(MultiPreOrderIteration(root))
	fmt.Println(MultiPreOrderRecursion(root))
}
