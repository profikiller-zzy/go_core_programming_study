package main

import "fmt"

/*
📝 N叉树标记移动问题
现状：
给定一个不规则的 N 叉树，节点总数为 p，把标记随机地放在这些节点上，使得标记总数也为 p。
——现在是不均匀分布。
目标：
我们现在需要挪动标记，使得每个节点上有且仅有1个标记。
——变成均匀分布。
限制条件：
标记一次只能挪动1个。
挪动仅支持相邻节点。例如父节点往子节点挪动，子节点往父节点挪动。
问题：
我们最少需要挪动多少次，能够达成「均匀分布」这个目标？
*/

type MultiTreeNode struct {
	Val      int
	Children []*MultiTreeNode
}

func NTtreeMarker(root *MultiTreeNode) int {
	stack := []*MultiTreeNode{root}
	balanceMap := make(map[*MultiTreeNode]int)
	var prev *MultiTreeNode
	var result int
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		if len(top.Children) == 0 || (prev != nil && top.Children[len(top.Children)-1] == prev) {
			stack = stack[:len(stack)-1]
			curBalance := top.Val - 1
			for _, child := range top.Children {
				curBalance += balanceMap[child]
			}
			balanceMap[top] = curBalance
			if top != root {
				result += abs(curBalance)
			}
			prev = top
		} else {
			// 倒序入栈
			for index := len(top.Children) - 1; index >= 0; index-- {
				stack = append(stack, top.Children[index])
			}
		}
	}
	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	leaf := &MultiTreeNode{
		Val: 0,
	}

	child := &MultiTreeNode{
		Val:      2,
		Children: []*MultiTreeNode{leaf},
	}

	root := &MultiTreeNode{
		Val:      1,
		Children: []*MultiTreeNode{child},
	}
	fmt.Println(NTtreeMarker(root))
}
