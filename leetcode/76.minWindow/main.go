package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	tByteCount, window := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tByteCount[t[i]]++
	}
	var start, validCount int
	var result string
	minLen := math.MaxInt32
	for end := 0; end < len(s); end++ {
		if _, ok := tByteCount[s[end]]; ok {
			window[s[end]]++
			if window[s[end]] == tByteCount[s[end]] {
				validCount++
			}
		}
		for validCount == len(tByteCount) {
			if end-start+1 < minLen {
				minLen = end - start + 1
				result = s[start : end+1]
			}
			if _, ok := tByteCount[s[start]]; ok {
				if window[s[start]] == tByteCount[s[start]] {
					validCount--
				}
				window[s[start]]--
			}
			start++
		}
	}
	return result
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t)) // Output: "BANC"
}
