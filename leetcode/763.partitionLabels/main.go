package main

import "fmt"

func partitionLabels(s string) []int {
	n := len(s)
	if n == 0 {
		return []int{}
	}

	var lastPos [26]int
	for i := 0; i < n; i++ {
		lastPos[s[i]-'a'] = i
	}
	res := make([]int, 0, 26)

	prevEnd := -1
	maxEnd := 0

	for i := 0; i < n; i++ {
		idx := s[i] - 'a'
		// 更新当前片段的最远边界
		if lastPos[idx] > maxEnd {
			maxEnd = lastPos[idx]
		}
		if i == maxEnd {
			res = append(res, i-prevEnd)
			prevEnd = i
		}
	}
	return res
}

func main() {
	s := "ababcbacadefegdehijhklij"
	result := partitionLabels(s)
	fmt.Println(result)
}
