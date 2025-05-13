package main

func partitionLabels(s string) []int {
	// 记录每个字符最后出现的位置
	maxDis := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		maxDis[s[i]] = i
	}

	var res []int
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		end = max(end, maxDis[s[i]])
		if i == end {
			res = append(res, end-start+1)
			start = i + 1
		}
	}
	return res
}
