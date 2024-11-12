package main

import (
	"fmt"
	"sort"
)

// 练习5.10： 重写topoSort函数，用map代替切片并移除对key的排序代码。验证结果的正确性（结果不唯一）。

// prereqs记录了每个课程的前置课程
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
}

func main() {
	for i, course := range topoSort1(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

// topoSort
func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

// topoSort1 对函数topoSort的改写
func topoSort1(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string][]string)
	// items是需要处理的先修课图
	visitAll = func(items map[string][]string) {
		preItems := make(map[string][]string, 0)
		for k, v := range items {
			if !seen[k] { // 这门课还没有被遍历过
				seen[k] = true
				for _, pres := range v { // 遍历这门课的所有先修课
					if !seen[k] { // 当前先修课没有被遍历过
						value, ok := m[pres] // 判断当前先修课是否存在自己的先修课
						if ok {              // 存在
							preItems[pres] = value // 相当与把这个点（这门先修课）和其对应的边全部加入pre_items中继续递归处理
							// 即深度优先遍历
						} else {
							preItems[pres] = make([]string, 0) // 不存在则结束观测
						}
					}
				}
			}
			visitAll(preItems)
			order = append(order, k)
		}
	}

	visitAll(m)
	return order
}
