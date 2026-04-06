package main

import "fmt"

type pair struct {
	variable string  // 除数
	value    float64 // 被除数等于多少倍的除数
}

// calcEquation 图加dfs
func calcEquation1(equations [][]string, values []float64, queries [][]string) []float64 {
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
		val := dfs1(graph, a, b, visited)
		res = append(res, val)
	}
	return res
}

func dfs1(graph map[string][]pair, cur, target string, visited map[string]bool) float64 {
	if cur == target {
		return 1.0
	}
	visited[cur] = true
	for _, nei := range graph[cur] {
		if visited[nei.variable] { // 该变量已经被查看过
			continue
		}
		res := dfs1(graph, nei.variable, target, visited)
		if res > 0 {
			return res * nei.value
		}
	}
	return -1.0
}

type node struct {
	varName  string
	multiple float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	n := len(equations)
	graph := make(map[string][]node)
	for i := 0; i < n; i++ {
		var1, var2 := equations[i][0], equations[i][1]
		graph[var1] = append(graph[var1], node{var2, values[i]})
		graph[var2] = append(graph[var2], node{var1, 1 / values[i]})
	}
	var result []float64
	for _, query := range queries {
		start, end := query[0], query[1]
		// query 中存在未定义的变量
		if _, ok := graph[start]; !ok {
			result = append(result, -1.0)
			continue
		}
		if _, ok := graph[end]; !ok {
			result = append(result, -1.0)
			continue
		}
		if start == end {
			result = append(result, 1.0)
			continue
		}
		isVisited := make(map[string]bool)
		isVisited[start] = true
		queue := []node{{start, 1.0}}
		isFind := false
		for len(queue) > 0 {
			curVar, curMultiple := queue[0].varName, queue[0].multiple
			queue = queue[1:]
			for _, e := range graph[curVar] {
				if !isVisited[e.varName] {
					isVisited[e.varName] = true
					if e.varName == end {
						result = append(result, curMultiple*e.multiple)
						isFind = true
						break
					}
					queue = append(queue, node{e.varName, curMultiple * e.multiple})
				}
			}
		}
		if !isFind {
			result = append(result, -1.0)
		}
	}
	return result
}

func main() {
	equations := [][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}}
	values := []float64{1.5, 2.5, 5.0}
	queries := [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}}
	fmt.Println(calcEquation(equations, values, queries))
}
