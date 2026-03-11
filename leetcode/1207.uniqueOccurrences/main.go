package main

func uniqueOccurrences(arr []int) bool {
	numCount := make(map[int]int)
	for _, num := range arr {
		numCount[num]++
	}
	countMap := make(map[int]bool)
	for _, count := range numCount {
		if !countMap[count] {
			countMap[count] = true
		} else {
			return false
		}
	}
	return true
}
