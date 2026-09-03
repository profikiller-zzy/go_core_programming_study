package main

import "strings"

type line struct {
	words      []string
	wordsCount int
	space      int // 剩余多少个空格
}

func fullJustify(words []string, maxWidth int) []string {
	var index int
	lines := make([]line, 0)
	for index < len(words) {
		curLine := line{
			words:      make([]string, 0),
			wordsCount: 0,
			space:      maxWidth,
		}
		// 贪心，加上这个单词之后，剩余的空位是否能够分割这些单词
		for index < len(words) && curLine.space-len(words[index]) >= curLine.wordsCount {
			curLine.words = append(curLine.words, words[index])
			curLine.wordsCount++
			curLine.space -= len(words[index])
			index++
		}
		lines = append(lines, curLine)
	}
	res := make([]string, len(lines))
	for index = 0; index < len(lines); index++ { // 最后一行不需要处理
		if index < len(lines)-1 {
			res[index] = lineToString(lines[index])
		} else {
			res[index] = lastLineToString(lines[index])
		}
	}
	return res
}

func lineToString(l line) string {
	if l.wordsCount == 1 {
		return l.words[0] + strings.Repeat(" ", l.space) // 如果没有单词，直接返回空格
	}

	var res strings.Builder
	var (
		partNum   int
		remainder int
		basic     int
	)
	partNum = l.wordsCount - 1    // 分为多少份
	remainder = l.space % partNum // 多出来的空格，需要均予加在前面的空格分隔中
	basic = l.space / partNum     // 每个分隔的空格数
	res.WriteString(l.words[0])
	for index := 1; index < len(l.words); index++ {
		if remainder > 0 {
			res.WriteString(strings.Repeat(" ", basic+1))
			remainder--
		} else {
			res.WriteString(strings.Repeat(" ", basic))
		}
		res.WriteString(l.words[index])
	}
	return res.String()
}

func lastLineToString(l line) string {
	var res strings.Builder
	res.WriteString(l.words[0])
	for index := 1; index < len(l.words); index++ {
		res.WriteString(" " + l.words[index])
		l.space--
	}
	if l.space > 0 {
		res.WriteString(strings.Repeat(" ", l.space))
	}
	return res.String()
}

func main() {
	words := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	maxWidth := 20
	result := fullJustify(words, maxWidth)
	for _, line := range result {
		println(line)
	}
}
