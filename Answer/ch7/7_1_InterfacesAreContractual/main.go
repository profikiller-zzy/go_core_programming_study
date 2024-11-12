package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

// 练习 7.1： 使用来自ByteCounter的思路，实现一个针对单词和行数的计数器。你会发现`bufio.ScanWords`非常的有用。

type wordCounter int
type lineCounter int

func (c *wordCounter) Write(p []byte) (int, error) {
	var scanner = bufio.NewScanner(bytes.NewReader(p))
	var words int
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*c++
		words++
	}
	return words, nil
}

func (c *lineCounter) Write(p []byte) (int, error) {
	var scanner = bufio.NewScanner(bytes.NewReader(p))
	var lines int
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		*c++
		lines++
	}
	return lines, nil
}

// 练习 7.2： 写一个带有如下函数签名的函数CountingWriter，传入一个io.Writer接口类型，
// 返回一个把原来的Writer封装在里面的新的Writer类型和一个表示新的写入字节数的int64类型指针。
//
// func CountingWriter(w io.Writer) (io.Writer, *int64)

// CountWriter 创建一个自定义的 Writer 类型，该类型嵌入了传入的 io.Writer 接口，并且包含一个 int64 类型的指针，
// 用于记录写入的字节数。每次写入时，将计算写入的字节数并更新指针的值
type CountWriter struct {
	writer io.Writer
	count  *int64
}

func (cw *CountWriter) Write(p []byte) (n int, err error) {
	n, err = cw.writer.Write(p)
	// 将写入的字节数累加至 *cw.count
	*cw.count += int64(n)
	return
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var count int64
	cw := &CountWriter{writer: w, count: &count}
	return cw, cw.count
}

// 练习 7.3： 为在gopl.io/ch4/treesort（§4.4）中的*tree类型实现一个String方法去展示tree类型的值序列。

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	if t == nil {
		return ""
	}
	var values []int
	values = appendValues(t, values) // 获取数节点值的中序序列

	var buf bytes.Buffer
	for i, v := range values {
		if i > 0 {
			fmt.Fprintf(&buf, " ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	return buf.String()
}

func appendValues(t *tree, values []int) []int {
	if t != nil { // 中序遍历获取节点序列
		values = appendValues(t.left, values)
		values = append(values, t.value)
		values = appendValues(t.right, values)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	// 递归添加树节点
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	//text := []byte("Hello, this is a sample text.\nIt has multiple lines.\nLet's count the words and lines.")
	//wc := new(wordCounter)
	////fmt.Fprintf(wc, string(text))
	//wc.Write(text)
	//lc := new(lineCounter)
	//lc.Write(text)
	//fmt.Println(int(*wc), int(*lc))

	root := &tree{value: 3}
	root.left = &tree{value: 1}
	root.right = &tree{value: 5}
	root.left.right = &tree{value: 2}
	root.right.left = &tree{value: 4}
	fmt.Println(root) // 输出: 1 2 3 4 5
}
