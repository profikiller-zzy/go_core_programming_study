package _5_8

import (
	"fmt"
	"golang.org/x/net/html"
)

// 练习 5.8： 修改pre和post函数，使其返回布尔类型的返回值。返回false时，中止forEachNoded的遍历。
// 使用修改后的代码编写ElementByID函数，根据用户输入的id查找第一个拥有该id元素的HTML元素，查找成功后，停止遍历。

func ElementByID(doc *html.Node, pre func(node *html.Node, string2 *string) bool, id string) *html.Node {
	// 当前节点就是要寻找的节点
	if doc != nil && pre(doc, &id) {
		return doc
	} else { // 遍历子节点和下一个兄弟节点
		if doc.FirstChild != nil {
			child := ElementByID(doc.FirstChild, pre, id)
			if child != nil {
				return child
			}
		}
		if doc.NextSibling != nil {
			sibling := ElementByID(doc.NextSibling, pre, id)
			if sibling != nil {
				return sibling
			}
		}
	}
	// 遍历完所有的节点后也没有找到
	return nil
}

func Pre(node *html.Node, id *string) bool {
	if node.Type == html.ElementNode && node.Data == *id {
		fmt.Println(node.Data)
		return true
	}
	return false
}
