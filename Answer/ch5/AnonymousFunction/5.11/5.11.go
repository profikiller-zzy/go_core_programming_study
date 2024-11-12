package main

import (
	"fmt"
)

// 练习5.11： 现在线性代数的老师把微积分设为了前置课程。完善topSort，使其能检测有向图中的环。

// prereqs 记录了每个课程的前置课程
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
	//"linear algebra":        {"calculus"},
}

func main() {
	fmt.Println(hasCycle(prereqs))
}

func hasCycle(m map[string][]string) bool {
	// visited 用于记录已经访问过的节点
	visited := make(map[string]bool, 0)
	var dfs func(string, map[string]bool) bool
	// item为当前节点，path记录路径
	dfs = func(item string, path map[string]bool) bool {
		if visited[item] == true { // 如果当前节点已经被访问过
			return false
		}
		// 将当前节点标记为已访问，并将其加入到path中
		visited[item] = true
		path[item] = true
		// 遍历所有当前节点的邻居节点
		for _, value := range m[item] {
			if path[value] || dfs(value, path) { // 如果有某个邻居节点在路径中，则说明有回路
				return true
			}
		}
		// 将当前节点从路径中删除
		delete(path, item)
		return false
	}

	// 遍历图中的所有节点
	for key, _ := range m {
		path := make(map[string]bool, 0)
		if dfs(key, path) {
			return true
		}
	}
	// 如果所有节点都检查完了没有环，则这个图没有环
	return false
}
