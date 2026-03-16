package main

import (
	"fmt"
	"sort"
)

type intervalType [][]int

func (i intervalType) Len() int {
	return len(i)
}

func (i intervalType) Less(j, k int) bool {
	return i[j][0] < i[k][0]
}

func (i intervalType) Swap(j, k int) {
	i[j], i[k] = i[k], i[j]
}

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	cur := []int{points[0][0], points[0][1]}
	count := 1
	for _, p := range points[1:] {
		if p[0] <= cur[1] { // 两个区间有重叠，则合并
			cur[0] = p[0] // 已经按照起点排序了,p[0]一定大于等于cur[0]
			cur[1] = min(cur[1], p[1])
		} else {
			count++
			cur = p
		}
	}
	return count
}

func main() {
	points := [][]int{
		{3, 9},
		{7, 12},
		{3, 8},
		{6, 8},
		{9, 10},
		{2, 9},
		{0, 9},
		{3, 9},
		{0, 6},
		{2, 8},
	}
	fmt.Println(findMinArrowShots(points))
}
