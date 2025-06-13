package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows <= 1 || numRows >= len(s) {
		return s
	}
	lines := make([][]byte, numRows)
	var index int
	for index = 0; index < len(s); {
		i := index
		nextPos := index + numRows
		// 处理从上到下的部分
		for i < len(s) && i < nextPos {
			lines[i-index] = append(lines[i-index], s[i])
			i++
		}
		nextPos = index + 2*numRows - 2
		// 处理从坐下到右上的部分
		for i < len(s) && i < nextPos {
			lines[nextPos-i] = append(lines[nextPos-i], s[i])
			i++
		}
		if i == len(s) {
			break
		} else {
			index = i
		}
	}
	var res strings.Builder
	for _, line := range lines {
		res.Write(line)
	}
	return res.String()
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 4))
}
