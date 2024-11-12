package _5_3

import "golang.org/x/net/html"

// 练习 5.3： 编写函数输出所有text结点的内容。注意不要访问<script>和<style>元素，因为这些元素对浏览者是不可见的。

func GetTextNodes(textNodes []string, node *html.Node) []string {
	if node == nil {
		return textNodes
	}
	if node.Type == html.ElementNode && (node.Data == "script" || node.Data == "style") {
		return textNodes
	}
	if node.Type == html.TextNode {
		textNodes = append(textNodes, node.Data)
	}
	// 递归处理子节点
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		textNodes = GetTextNodes(textNodes, c)
	}
	return textNodes
}
