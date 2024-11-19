package main

import (
	"fmt"
	"io"
	"strings"
)

// 练习 7.4： strings.NewReader函数通过读取一个string参数返回一个满足io.Reader接口类型的值（和其它值）。
// 实现一个简单版本的NewReader，用它来构造一个接收字符串输入的HTML解析器（§5.2）

type HtmlParser struct {
	html  []byte // html文本的字节切片
	index int    // 读的位置
}

// NewReader 返回一个满足 io.Reader 接口的值，从给定的字符串开始读取
func NewReader(html string) *HtmlParser {
	return &HtmlParser{html: []byte(html), index: 0}
}

// Read 实现了 io.Reader 接口的 Read 方法
func (hp *HtmlParser) Read(b []byte) (n int, err error) {
	if hp.index >= len(hp.html) {
		return 0, io.EOF // 已经读取到字符串的末尾，返回 EOF
	}

	n = copy(b, hp.html[hp.index:]) // 将剩余的字符串拷贝到 b 中
	hp.index += n

	return n, nil
}

// 练习 7.5： io包里面的LimitReader函数接收一个io.Reader接口类型的r和字节数n，
// 并且返回另一个从r中读取字节但是当读完n个字节后就表示读到文件结束的Reader。实现这个LimitReader函数：
// func LimitReader(r io.Reader, n int64) io.Reader

type Reader struct {
	index int // 读取的位置
	read  io.Reader
	n     int // 被限制读取的字节数，一旦读取到n个字节则返回EOF
}

// Read 实现这个被限制的Reader
func (r *Reader) Read(buf []byte) (n int, err error) {
	// 如果已经读了n个字节，则直接返回EOF
	if r.index >= r.n {
		return 0, io.EOF
	}
	// 计算可以读取的字节数
	remaining := r.n - r.index
	if len(buf) > remaining {
		buf = buf[:remaining] // 缓冲区大小限制为剩余可读字节数
	}

	// 从嵌套的 reader 中读取数据
	n, err = r.read.Read(buf)

	// 更新读取的字节数索引
	r.index += n

	// 如果读取到数据但到达了限制，返回 io.EOF
	if r.index >= r.n {
		err = io.EOF
	}

	return n, err
}

// LimitReader 获取一个限制只能读取N个字节的io.Reader
func LimitReader(r io.Reader, n int) io.Reader {
	return &Reader{index: 0, read: r, n: n}
}

func main() {
	//s := "<html><head><title>Hello</title></head><body><h1>Hello, World!</h1></body></html>"
	//reader := NewReader(s)
	//
	//buf := make([]byte, 32) // 缓冲区
	//for {
	//	n, err := reader.Read(buf)
	//	if err != nil {
	//		if err == io.EOF {
	//			break // 读取完毕，退出循环
	//		}
	//		// 处理其他可能的错误
	//		break
	//	}
	//	// 处理读取的数据
	//	println(string(buf[:n]))
	//}

	// 创建一个读入流
	r := strings.NewReader("1234567890")
	r1 := LimitReader(r, 4)
	//s := bufio.NewScanner(r1)
	//s.Split(bufio.ScanBytes)
	//for s.Scan() {
	//	fmt.Println(s.Text())
	//}
	var buf []byte
	buf = make([]byte, 8)
	n, err := r1.Read(buf)
	fmt.Println(string(buf))
	fmt.Printf("n: %d\t err: %v\n", n, err)
}
