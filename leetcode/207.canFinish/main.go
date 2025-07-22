package main

func canFinish1(numCourses int, prerequisites [][]int) bool {
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

func canFinish(numCourses int, prerequisites [][]int) bool {
	inDegree := map[int]int{}
	graph := make([][]int, numCourses)
	queue := make([]int, 0) // 队列用于存储入度为0的节点
	for index := 0; index < numCourses; index++ {
		graph[index] = make([]int, 0) // 使用邻接表存储图
		inDegree[index] = 0
	}
	for _, pre := range prerequisites {
		inDegree[pre[0]]++
		graph[pre[1]] = append(graph[pre[1]], pre[0])
	}
	for course, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, course)
		}
	}
	var total int
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		total++
		for _, v := range graph[cur] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	if total == numCourses {
		return true
	}
	return false
}

func main() {

}
