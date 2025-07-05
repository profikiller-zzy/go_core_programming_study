package main

func findOrder1(numCourses int, prerequisites [][]int) []int {
	inDegree := make([]int, numCourses) // 入度数组
	graph := make([][]int, numCourses)  // 邻接表
	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, numCourses)
	}
	for _, pre := range prerequisites {
		inDegree[pre[0]]++
		graph[pre[1]][pre[0]] = 1
	}

	var result []int
	for {
		// 找到入度为0的节点
		found := false
		for i := 0; i < numCourses; i++ {
			if inDegree[i] == 0 {
				found = true
				inDegree[i] = -1 // 标记为已访问
				result = append(result, i)
				for j := 0; j < numCourses; j++ {
					if graph[i][j] == 1 {
						inDegree[j]--
					}
				}
			}
		}
		if !found {
			break
		}
	}

	for i := 0; i < numCourses; i++ {
		if inDegree[i] != -1 {
			return nil // 有环
		}
	}
	return result
}

// findOrder 使用拓扑排序+队列+邻接表解决
func findOrder(numCourses int, prerequisites [][]int) []int {
	inDegree := make([]int, numCourses)
	graph := make([][]int, numCourses)

	// 构建邻接表和入度数组
	for _, pre := range prerequisites {
		from, to := pre[1], pre[0]
		graph[from] = append(graph[from], to)
		inDegree[to]++
	}

	// 入度为0的点入队
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	result := make([]int, 0)

	// BFS 拓扑排序
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		result = append(result, curr)

		for _, neighbor := range graph[curr] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// 如果结果中课程数量不等于总课程数，则说明存在环
	if len(result) != numCourses {
		return nil
	}

	return result
}
