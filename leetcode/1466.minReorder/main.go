package main

func minReorder(n int, connections [][]int) int {
	// 先建图，这个图和邻接矩阵不同
	edges := make([][][2]int, n)
	for _, conn := range connections {
		from, to := conn[0], conn[1]
		// 从from看，从from遍历到to时需要反向，于是设置为1
		edges[from] = append(edges[from], [2]int{to, 1})
		// 从to看，从to遍历到from时不需要反向，于是设置为0
		edges[to] = append(edges[to], [2]int{from, 0})
	}

	isVisited := make([]bool, n)
	var count int
	var dfs func(int)
	// 当我们遍历的时候，肯定是从0开始遍历这个树，走完所有的边
	// 对于正在遍历的边，这条边如果是到达当前节点的则不用反向，反之则需要，就将count+1
	dfs = func(city int) {
		isVisited[city] = true
		for _, edge := range edges[city] {
			nextCity := edge[0]
			if !isVisited[nextCity] {
				count += edge[1]
				dfs(nextCity)
			}
		}
	}
	dfs(0)
	return count
}
