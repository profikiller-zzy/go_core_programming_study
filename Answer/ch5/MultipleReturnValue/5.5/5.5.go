package _5_5

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

// 练习 5.5： 实现countWordsAndImages。（参考练习4.9如何分词）

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

// countWordsAndImages 计算单词和图片的数量
func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}
	if n.Type == html.TextNode {
		wordList := strings.Split(strings.TrimSpace(n.Data), " ")
		words += len(wordList)
	}
	// 如果当前节点不是 <script> 或 <style> 标签，则递归遍历子节点
	if n.Type == html.ElementNode && (n.Data != "script" && n.Data != "style") {
		w, i := countWordsAndImages(n.FirstChild)
		words += w
		images += i
	}
	if n.NextSibling != nil {
		ws, is := countWordsAndImages(n.NextSibling)
		words += ws
		images += is
	}
	// 如果当前节点是 <img> 标签，则增加图片计数
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	return words, images
}
