package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	words := strings.Split(strings.TrimSpace(s), " ")
	return len(words[len(words)-1])
}

func main() {
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  ")) // Output: 5
}
