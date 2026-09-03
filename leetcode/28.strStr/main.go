package main

import "fmt"

func strStr(haystack string, needle string) int {
	var sIndex, pIndex int
	next := buildNext(needle)
	for sIndex < len(haystack) {
		if haystack[sIndex] == needle[pIndex] {
			sIndex++
			pIndex++
			if pIndex == len(needle) {
				return sIndex - pIndex
			}
			continue
		}
		fail := next[pIndex]
		matched := false
		for fail != -1 {
			if needle[fail+1] == haystack[sIndex] {
				pIndex = fail + 2
				matched = true
				break
			}
			fail = next[fail]
		}
		if !matched && needle[0] == haystack[sIndex] {
			pIndex = 1
			matched = true
		}
		if !matched {
			pIndex = 0
		}
		sIndex++
	}
	return -1
}

func buildNext(pattern string) []int {
	n := len(pattern)
	next := make([]int, n)
	next[0] = -1
	i, j := 0, -1
	for i <= n-1 {
		if j == -1 || pattern[i] == pattern[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}

func main() {
	s := "mississippi"
	pattern := "issip"
	fmt.Println(strStr(s, pattern))
}
