package main

import "fmt"

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
	equations := [][]string{{"a", "b"}, {"b", "c"}}
	values := []float64{2.0, 3.0}
	queries := [][]string{{"a", "c"}, {"b", "a"}}
	result := calcEquation(equations, values, queries)
	fmt.Println(result)
}
