package main

import (
	"fmt"
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {
	var strBuild strings.Builder
	i, j := 0, 0
	for i < len(word1) || j < len(word2) {
		if i < len(word1) {
			strBuild.WriteByte(word1[i])
			i++
		}
		if j < len(word2) {
			strBuild.WriteByte(word2[j])
			j++
		}
	}
	return strBuild.String()
}

func main() {
	word1 := "ab"
	word2 := "pqr"
	fmt.Println(mergeAlternately(word1, word2))
}
