package main

import (
	"strings"
)

func convert1(s string, numRows int) string {
	length := len(s)
	if numRows == 1 || length < numRows {
		return s
	}
	// 先计算需要用到几列
	numCols := (length - 1) * (length / (2 * (numRows - 1)))
	sur := length % (2 * (length - 1))
	if sur > 0 {
		if sur > numRows {
			numCols += sur - numRows + 1
		} else {
			numCols += 1
		}
	}

	arr := make([][]rune, numRows)
	for i := 0; i < numRows; i++ {
		arr[i] = make([]rune, numCols)
	}

	index := 0
	for col := 0; col < numCols && index < length; col++ {
		if col%(numRows-1) == 0 {
			for i := 0; i < numRows; i++ {
				arr[i][col] = rune(s[index])
				index++
				if index >= length {
					break
				}
			}
		} else {
			arr[numRows-1-col%(numRows-1)][col] = rune(s[index])
			index++
		}
	}

	var builder strings.Builder

	for _, row := range arr {
		for _, ch := range row {
			if ch == 0 {
				builder.WriteString(" ")
			} else {
				builder.WriteRune(ch)
			}
		}
		builder.WriteString("\n")
	}
	// 返回构建好的字符串
	return builder.String()
}

func convert(s string, numRows int) string {
	// 如果只有一行，或者字符串的长度小于 numRows，直接返回原始字符串
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	// 创建一个切片，每个切片存储 Z 字形的一行
	rows := make([]strings.Builder, numRows)
	curRow := 0
	goingDown := false

	// 遍历字符串，将每个字符放到对应的行
	for _, char := range s {
		rows[curRow].WriteRune(char)

		// 当到达最顶或最底行时，改变方向
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}

		// 根据方向移动行索引
		if goingDown {
			curRow++
		} else {
			curRow--
		}
	}

	// 将所有行拼接成一个最终的字符串
	var result strings.Builder
	for _, row := range rows {
		result.WriteString(row.String())
	}

	return result.String()
}

func main() {
	s := "PAYPALISHIRING"
	numRows := 3
	result := convert(s, numRows)
	println(result) // 输出: "PAHNAPLSIIGYIR"
}
