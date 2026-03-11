package main

import (
	"fmt"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func postorder(root *TreeNode) []int {
	var stack []*TreeNode
	result := make([]int, 0)
	if root == nil {
		return result
	}
	cur := root
	var prev *TreeNode
	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.left
		}
		top := stack[len(stack)-1]
		if top.right == nil || top.right == prev {
			result = append(result, top.val)
			stack = stack[:len(stack)-1]
			prev = top
		} else {
			cur = top.right
		}
	}
	return result
}

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	s2 := s1[2:6:7]

	s2 = append(s2, 100)
	s2 = append(s2, 200)

	s1[2] = 20

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(slice)
}
