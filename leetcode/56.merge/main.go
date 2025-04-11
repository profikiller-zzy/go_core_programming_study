package main

import "sort"

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

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 排序
	inv := intervalType(intervals)
	sort.Sort(inv)

	result := [][]int{inv[0]}
	for i := 1; i < len(inv); i++ {
		last := result[len(result)-1]
		curr := inv[i]
		if last[1] >= curr[0] {
			// 合并区间，更新右边界
			if curr[1] > last[1] {
				last[1] = curr[1]
			}
		} else {
			// 不重叠，添加到结果中
			result = append(result, curr)
		}
	}
	return result
}

func main() {

}
