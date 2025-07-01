package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	curPtr *TreeNode
	head   *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	var prev *TreeNode
	cur, head := root, root
	for head.Left != nil {
		head = head.Left
	}
	for cur != nil {
		if cur.Left == nil {
			prev = cur
			cur = cur.Right
			continue
		}
		leftRoot, leftPrev := cur.Left, cur.Left
		for leftPrev.Right != nil {
			leftPrev = leftPrev.Right
		}
		cur.Left = nil
		leftPrev.Right = cur
		cur = leftRoot
		if prev != nil {
			prev.Right = cur
		}
	}
	return BSTIterator{
		curPtr: nil,
		head:   head,
	}
}

func (this *BSTIterator) Next() int {
	if this.curPtr == nil {
		this.curPtr = this.head
		return this.curPtr.Val
	}
	this.curPtr = this.curPtr.Right
	return this.curPtr.Val
}

func (this *BSTIterator) HasNext() bool {
	if this.curPtr == nil {
		return this.head != nil
	}
	return this.curPtr.Right != nil
}
