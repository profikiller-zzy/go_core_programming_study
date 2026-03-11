package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	dummy := &TreeNode{Left: root}
	cur, curParent := root, dummy
	for cur != nil {
		if key == cur.Val {
			break
		}
		curParent = cur
		if key > cur.Val {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}

	// 没有找到这个key
	if cur == nil {
		return dummy.Left
	}

	// 执行删除
	if cur.Left == nil { // 用右节点代替cur
		if curParent.Left == cur {
			curParent.Left = cur.Right
		} else {
			curParent.Right = cur.Right
		}
	} else if cur.Right == nil { // 用左节点代替
		if curParent.Left == cur {
			curParent.Left = cur.Left
		} else {
			curParent.Right = cur.Left
		}
	} else { // cur 左右节点都存在，就需要找到中序遍历中cur的后继节点来代替cur
		successorParent := cur
		successor := cur.Right
		for successor.Left != nil {
			successorParent = successor
			successor = successor.Left
		}
		cur.Val = successor.Val
		if successorParent.Left == successor { // 后继节点顶多只有右节点
			successorParent.Left = successor.Right
		} else { // 后继节点是cur.Right的直接情况
			successorParent.Right = successor.Right
		}
	}
	return dummy.Left
}
