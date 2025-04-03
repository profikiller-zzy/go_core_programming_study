package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/count-and-say/submissions/618203992/
func countAndSay(n int) string {
	prevText := "1"
	for index := 0; index < n-1; index++ {
		curText := strings.Builder{}
		for textIndex, pos := 0, 0; textIndex < len(prevText); pos = textIndex {
			for textIndex < len(prevText) && prevText[textIndex] == prevText[pos] { // 统计相同字符的个数
				textIndex++
			}
			curText.WriteString(strconv.Itoa(textIndex - pos))
			curText.WriteByte(prevText[pos])
		}
		prevText = curText.String()
	}
	return prevText
}

func main() {
	fmt.Println(countAndSay(5))
}
