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

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	type edge struct {
		to       string
		multiple float64
	}
	edges := make(map[string][]edge)
	for i, equa := range equations {
		from, to := equa[0], equa[1]
		edges[from] = append(edges[from], edge{to, values[i]})
		edges[to] = append(edges[to], edge{from, 1 / values[i]})
	}

	var dfs func(string, string, float64, map[string]bool) float64
	dfs = func(cur, target string, curMul float64, isVisited map[string]bool) float64 {
		isVisited[cur] = true
		if cur == target {
			return curMul
		}
		for _, e := range edges[cur] {
			to := e.to
			multiple := e.multiple
			if !isVisited[to] {
				result := dfs(to, target, curMul*multiple, isVisited)
				if result != -1 {
					return result
				}
			}
		}
		return -1.0
	}
	result := make([]float64, 0, len(queries))
	for _, query := range queries {
		var1, var2 := query[0], query[1]
		if _, ok := edges[var1]; !ok {
			result = append(result, -1.0)
			continue
		}
		if _, ok := edges[var2]; !ok {
			result = append(result, -1.0)
			continue
		}
		isVisited := make(map[string]bool)
		queryVal := dfs(var1, var2, 1.0, isVisited)
		result = append(result, queryVal)
	}
	return result
}

func main() {
	equations := [][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}}
	values := []float64{1.5, 2.5, 5.0}
	queries := [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}}
	fmt.Println(calcEquation(equations, values, queries))
}
