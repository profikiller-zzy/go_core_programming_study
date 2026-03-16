package main

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals1(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// 按起点排序，起点相同按终点排序
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	var count int
	prev := intervals[0]

	for _, inv := range intervals[1:] {
		if inv[0] < prev[1] { // 重叠
			count++
			// 保留终点更小的区间
			if inv[1] < prev[1] {
				prev = inv
			}
		} else {
			prev = inv
		}
	}

	return count
}

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// 按终点排序，起点相同按终点排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	// 要计算最少需要删除的区间，我们可以计算最大不重叠的区间总数，用总区间数减去这个数就是结果
	count, right := 1, intervals[0][1]
	for _, inv := range intervals[1:] {
		// 由于我们是按照终点排序的，所以当前区间的终点一定小于等于后面区间的终点
		if inv[0] >= right {
			count++
			right = inv[1]
		}
	}
	return len(intervals) - count
}

func main() {
	intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	fmt.Println(eraseOverlapIntervals(intervals))
}
