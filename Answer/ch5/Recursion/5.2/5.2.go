package _5_2

import (
	"golang.org/x/net/html"
)

// 练习 5.2： 编写函数，记录在HTML树中出现的同名元素的次数。

// record  记录在HTML树中出现的同名元素的次数
func Record(count map[string]int, node *html.Node) {
	// 当前节点为空直接返回
	if node == nil {
		return
	}
	if node.Type == html.ElementNode {
		count[node.Data]++
	}
	// 遍历子节点
	Record(count, node.FirstChild)
	// 遍历下一个兄弟节点
	Record(count, node.NextSibling)
}
