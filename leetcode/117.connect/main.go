package main

// Node Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	curLevel := []*Node{root}
	for len(curLevel) > 0 {
		nextLevel := make([]*Node, 0)
		for i, node := range curLevel {
			if i < len(curLevel)-1 {
				node.Next = curLevel[i+1]
			} else {
				node.Next = nil
			}
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		curLevel = nextLevel
	}
	return root
}
