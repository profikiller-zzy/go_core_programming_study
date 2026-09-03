package main

import "fmt"

func buildNext(pattern string) []int {
	next := make([]int, len(pattern))
	next[0] = -1
	i, j := 0, -1
	for i < len(pattern)-1 {
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

func KMPSearch(text, pattern string) int {
	if len(pattern) == 0 {
		return 0
	}
	if len(pattern) > len(text) {
		return -1
	}

	next := buildNext(pattern)
	textIndex, patternIndex := 0, 0

	for textIndex < len(text) {
		if text[textIndex] == pattern[patternIndex] {
			textIndex++
			patternIndex++
			if patternIndex == len(pattern) {
				return textIndex - patternIndex
			}
			continue
		}

		if patternIndex != 0 {
			patternIndex = next[patternIndex]
			continue
		}

		textIndex++
	}

	return -1
}

func main() {
	text := "ababcabcacbab"
	pattern := "abcac"

	fmt.Printf("text=%q, pattern=%q\n", text, pattern)
	fmt.Printf("first match index: %d\n", KMPSearch(text, pattern))
}
