package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)
	for left, right := 0, len(words)-1; left < right; {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}
	var res strings.Builder
	res.WriteString(words[0])
	for index := 1; index < len(words); index++ {
		res.WriteString(" " + words[index])
	}
	return res.String()
}

func main() {
	fmt.Println(reverseWords("a good   example"))
}
