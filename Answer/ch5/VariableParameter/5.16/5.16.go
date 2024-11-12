package main

import (
	"fmt"
	"strings"
)

// 练习5.16： 编写多参数版本的strings.Join。

// Join 函数接受一个分隔符 sep 和一个可变长度的字符串元素 elems，
// 并将 elems 中的元素用 sep 分隔符连接起来，返回连接后的字符串。
func Join(sep string, elems ...string) string {
	// 处理特殊情况。
	switch len(elems) {
	case 0:
		// 如果 elems 为空，返回空字符串。
		return ""
	case 1:
		// 如果 elems 只有一个元素，返回该元素。
		return elems[0]
	}

	// 计算连接后的字符串长度。
	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		n += len(elems[i])
	}

	// 使用 strings.Builder 来高效地构建连接后的字符串。
	var b strings.Builder
	b.Grow(n)
	b.WriteString(elems[0])
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(s)
	}

	// 返回连接后的字符串。
	return b.String()
}

func main() {
	// 不太懂这个题的意思
	// 是将 elems 切片参数改为可变参数吗？
	str := []string{"a", "b", "c", "d", "e", "f"}
	// str... 为解切片操作。
	fmt.Println(Join("-", str...))
	fmt.Println(Join(",", "abc", "def", "ghi"))
}
