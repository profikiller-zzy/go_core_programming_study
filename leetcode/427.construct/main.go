package main

// Node Definition for a QuadTree node
type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

type pos struct {
	row, col int
}

func construct(grid [][]int) *Node {
	var dfs func(topLeft, bottomRight pos) *Node
	dfs = func(topLeft, bottomRight pos) *Node {
		if topLeft.row == bottomRight.row && topLeft.col == bottomRight.col {
			return &Node{
				Val:    grid[topLeft.row][topLeft.col] == 1,
				IsLeaf: true,
			}
		}
		midRow := (topLeft.row + bottomRight.row) / 2
		midCol := (topLeft.col + bottomRight.col) / 2
		topLeftNode := dfs(pos{topLeft.row, topLeft.col}, pos{midRow, midCol})
		topRightNode := dfs(pos{topLeft.row, midCol + 1}, pos{midRow, bottomRight.col})
		bottomLeftNode := dfs(pos{midRow + 1, topLeft.col}, pos{bottomRight.row, midCol})
		bottomRightNode := dfs(pos{midRow + 1, midCol + 1}, pos{bottomRight.row, bottomRight.col})
		if topLeftNode.IsLeaf && topRightNode.IsLeaf && bottomLeftNode.IsLeaf && bottomRightNode.IsLeaf &&
			topLeftNode.Val == topRightNode.Val && topLeftNode.Val == bottomLeftNode.Val && topLeftNode.Val == bottomRightNode.Val {
			return &Node{
				Val:    topLeftNode.Val,
				IsLeaf: true,
			}
		} else {
			return &Node{
				Val:         false,
				IsLeaf:      false,
				TopLeft:     topLeftNode,
				TopRight:    topRightNode,
				BottomLeft:  bottomLeftNode,
				BottomRight: bottomRightNode,
			}
		}
	}
	rows, cols := len(grid), len(grid[0])
	if rows == 0 || cols == 0 {
		return nil
	}
	return dfs(pos{0, 0}, pos{rows - 1, cols - 1})
}
