package main

func insert(intervals [][]int, newInterval []int) [][]int {
	var insetIndex int
	for insetIndex < len(intervals) && intervals[insetIndex][0] <= newInterval[0] {
		insetIndex++
	}
	intervals = append(intervals[:insetIndex], append([][]int{newInterval}, intervals[insetIndex:]...)...)

	result := [][]int{intervals[0]}
	for index := 1; index < len(intervals); index++ {
		last := result[len(result)-1]
		newCur := []int{intervals[index][0], intervals[index][1]}
		if last[1] >= newCur[0] { // 区间重叠
			if newCur[1] > last[1] {
				last[1] = newCur[1]
			}
		} else { // 区间不重叠
			result = append(result, newCur)
		}
	}
	return result
}
