package _5_7

import (
	"fmt"
	"golang.org/x/net/html"
)

// 练习 5.7： 完善startElement和endElement函数，使其成为通用的HTML输出器。
// 要求：输出注释结点，文本结点以及每个元素的属性（< a href='...'>）。
// 使用简略格式输出没有孩子结点的元素（即用<img/>代替<img></img>）。
// 编写测试，验证程序输出的格式正确。（详见11章）

var depth int

func StartElement(n *html.Node) {
	// getAttrString 函数值，用于将[]html.Attribute转换为对应的string
	var getAttrString func(attribute []html.Attribute) string
	getAttrString = func(attribute []html.Attribute) string {
		var res string
		for _, value := range n.Attr {
			res += fmt.Sprintf(" %s='%s'", value.Key, value.Val)
		}
		return res
	}
	// tag 为其实标签对应的string
	var tag string
	if n.Type == html.ElementNode {
		var attrString string
		attrString = getAttrString(n.Attr)
		tag = fmt.Sprintf("%*s<%s%s>\n", depth*2, "", n.Data, attrString)
		depth++
	}
	fmt.Printf(tag)
}
func EndElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

// ForEachNode 针对每个结点x，都会调用pre(x)和post(x)。
// pre和post都是可选的。
// 遍历孩子结点之前，pre被调用
// 遍历孩子结点之后，post被调用
func ForEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ForEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
