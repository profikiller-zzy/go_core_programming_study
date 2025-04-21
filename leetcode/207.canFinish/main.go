package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	inDegree := make([]int, numCourses) // 入度数组
	graph := make([][]int, numCourses)  // 邻接表
	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, numCourses)
	}
	for _, pre := range prerequisites {
		inDegree[pre[0]]++
		graph[pre[1]][pre[0]] = 1
	}

	for {
		// 找到入度为0的节点
		found := false
		for i := 0; i < numCourses; i++ {
			if inDegree[i] == 0 {
				found = true
				inDegree[i] = -1 // 标记为已访问
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
			return false // 有环
		}
	}
	return true
}

func main() {

}
