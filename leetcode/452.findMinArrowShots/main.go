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
	if len(points) <= 1 {
		return len(points)
	}
	inv := intervalType(points)
	sort.Sort(inv)

	result := [][]int{[]int{points[0][0], points[0][1]}}
	for index := 1; index < len(inv); index++ {
		cur := result[len(result)-1]
		next := inv[index]
		if cur[1] >= next[0] { // 两个区间是相交的
			cur[0] = next[0]
			cur[1] = min(cur[1], next[1])
		} else { // 两个区间不相交
			newPoint := []int{next[0], next[1]}
			result = append(result, newPoint)
		}
	}
	return len(result)
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
