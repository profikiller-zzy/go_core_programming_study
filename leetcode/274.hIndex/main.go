package main

import "sort"

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})
	var h int
	for i := 0; i < len(citations) && citations[i] >= h+1; i++ {
		h++
	}
	return h
}
