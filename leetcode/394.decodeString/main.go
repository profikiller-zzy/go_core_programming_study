package main

import (
	"fmt"
	"strings"
)

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func decodeString(s string) string {
	var (
		strStack []string
		numStack []int
		currStr  strings.Builder // 用于构建当前字符串，即题目中的 encoded_string
		curNum   int             // 用于解析当前数字
	)

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if isDigit(ch) {
			curNum = curNum*10 + int(ch-'0') // 支持多位数字
		} else if ch == '[' {
			strStack = append(strStack, currStr.String())
			numStack = append(numStack, curNum)
			currStr.Reset()
			curNum = 0
		} else if ch == ']' {
			lastStr := strStack[len(strStack)-1]
			curNum = numStack[len(numStack)-1]
			strStack = strStack[:len(strStack)-1]
			numStack = numStack[:len(numStack)-1]

			// 构建当前字符串 str = lastStr + currStr * curNum
			var temp strings.Builder
			temp.WriteString(lastStr)
			for j := 0; j < curNum; j++ {
				temp.WriteString(currStr.String())
			}
			curNum = 0 // 重置 curNum
			currStr.Reset()
			currStr.WriteString(temp.String())
		} else {
			currStr.WriteByte(ch)
		}
	}
	return currStr.String()
}

func main() {
	s := "3[a2[c]]abc"
	fmt.Println(decodeString(s))
}
