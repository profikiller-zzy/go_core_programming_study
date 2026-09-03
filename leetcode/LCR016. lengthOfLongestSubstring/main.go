package main

func lengthOfLongestSubstring(s string) int {
	pos := make(map[byte]int)
	start, end := 0, 0
	var res int
	for ; end < len(s); end++ {
		if p, ok := pos[s[end]]; ok {
			if p > start {
				start = p
			}
			res = max(res, end-start)
		}
		pos[s[end]] = end
	}
	res = max(res, end-start)
	return res
}
