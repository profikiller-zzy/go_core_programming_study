package _5_4

import "golang.org/x/net/html"

// 练习 5.4： 扩展visit函数，使其能够处理其他类型的结点，如images、scripts和style sheets。

func Visit(links []string, n *html.Node) []string {
	// 假如当前节点为空，直接结束
	if n == nil {
		return links
	}
	// 首先判断该节点是不是元素节点
	if n.Type == html.ElementNode {
		if n.Data == "link" || n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		} else if n.Data == "img" || n.Data == "script" {
			for _, a := range n.Attr {
				if a.Key == "src" {
					links = append(links, a.Val)
				}
			}
		}
	}
	// 这里是源代码循环调用的部分
	//for c := n.FirstChild; c != nil; c = c.NextSibling {
	//	links = visit(links, c)
	//}
	// 修改后，改为递归调用
	// 遍历子节点
	links = Visit(links, n.FirstChild)
	// 遍历下一个兄弟节点
	links = Visit(links, n.NextSibling)
	return links
}
