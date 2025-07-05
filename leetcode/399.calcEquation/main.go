package main

type pair struct {
	variable string  // 除数
	value    float64 // 被除数等于多少倍的除数
}

// calcEquation 图加dfs
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string][]pair) // key 为变量名，value是该变量作为被除数

	// 建图
	for i := range equations {
		a, b := equations[i][0], equations[i][1]
		k := values[i]
		graph[a] = append(graph[a], pair{b, k})
		graph[b] = append(graph[b], pair{a, 1 / k})
	}

	// 查询
	var res []float64
	for _, query := range queries {
		a, b := query[0], query[1]
		if _, ok := graph[a]; !ok || graph[b] == nil { // query 中有变量不存在
			res = append(res, -1.0)
			continue
		}
		visited := make(map[string]bool)
		val := dfs(graph, a, b, visited)
		res = append(res, val)
	}
	return res
}

func dfs(graph map[string][]pair, cur, target string, visited map[string]bool) float64 {
	if cur == target {
		return 1.0
	}
	visited[cur] = true
	for _, nei := range graph[cur] {
		if visited[nei.variable] { // 该变量已经被查看过
			continue
		}
		res := dfs(graph, nei.variable, target, visited)
		if res > 0 {
			return res * nei.value
		}
	}
	return -1.0
}
