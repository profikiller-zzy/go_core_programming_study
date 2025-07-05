package main

// Node Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	// 创建两个map，一个map存储node->newNode的映射关系，另一个map存储node->是否被访问过的映射关系
	mapping := make(map[*Node]*Node)
	visited := make(map[*Node]bool)

	var dfs func(*Node) *Node
	dfs = func(origin *Node) *Node {
		if visited[origin] {
			return mapping[origin] // 如果已经访问过，直接返回对应的新节点
		}
		visited[origin] = true                      // 标记当前节点已访问
		newNode := &Node{Val: origin.Val}           // 创建新节点
		mapping[origin] = newNode                   // 将原节点映射到新节点
		for _, neighbor := range origin.Neighbors { // 遍历邻居
			newNode.Neighbors = append(newNode.Neighbors, dfs(neighbor)) // 递归克隆邻居节点
		}
		return newNode
	}
	return dfs(node) // 从起始节点开始克隆图
}
