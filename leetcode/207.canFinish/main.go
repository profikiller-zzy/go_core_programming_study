package main

func canFinish1(numCourses int, prerequisites [][]int) bool {
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
	return total == numCourses
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	sign := make([]int, numCourses)
	for index := 0; index < numCourses; index++ {
		graph[index] = make([]int, 0)
	}
	var dfs func(int)
	valid := true
	stack := make([]int, 0, numCourses)
	dfs = func(course int) {
		sign[course] = 1
		for _, v := range graph[course] {
			if sign[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if sign[v] == 1 {
				valid = false
				return
			}
			sign[v] = 2
			stack = append(stack, course)
		}
	}
	for index := 0; index < numCourses; index++ {
		if sign[index] == 0 {
			dfs(index)
		}
	}
	return valid
}

func main() {

}
