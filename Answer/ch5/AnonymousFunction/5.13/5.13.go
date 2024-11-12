package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// 练习5.13： 修改crawl，使其能保存发现的页面，必要时，可以创建目录来保存这些页面。
// 只保存来自原始域名下的页面。假设初始页面在golang.org下，就不要保存vimeo.com下的页面。

// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool) // 用于记录已经处理过的元素
	for len(worklist) > 0 {
		// items是当前队列
		items := worklist
		// worklist存储下一轮要处理的页面URL
		worklist = nil
		// 遍历当前队列
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				// f(item)... 表示将string切片解包为可变参数，使`worklist`可以添加`f(item)`中返回值的所有元素
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

// crawl 返回当前页面所引用的所有链接
func crawl(url string) []string {
	fmt.Println(url)
	list, err := ExtractRestructure(url, true)
	if err != nil {
		log.Print(err)
	}
	return list
}

func ExtractRestructure(url string, isSave bool) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string

	// originURL 原始域名
	originURL := resp.Request.URL.Host
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				// 解析HTML中`<a>`标签中的`href`属性，将其解析为绝对地址
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
				// 如果不需要保存当前页面
				if isSave && link.Host != originURL {
					continue
				}
				// 保存页面
				// 1.创建文件
				filePath := getFilePath(a.Val)
				file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
				if err != nil { // 直接处理文件打开错误
					fmt.Printf("file %s open failed: %s", filePath, err)
				}
				//及时关闭file句柄
				defer file.Close()
				//写入文件时，使用带缓存的 *Writer
				write := bufio.NewWriter(file)
				if _, err := write.WriteString(link.String() + "\r\n"); err != nil {
					panic(err)
				}
				// 将缓存写入文件
				if err := write.Flush(); err != nil {
					panic(err)
				}
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func getFilePath(relativePath string) string {
	fileRoot, err := os.Getwd()
	if err != nil {
		fmt.Printf("get file root failed: %s\n", err)
		return ""
	}
	// 使用 filepath.Join 拼接路径
	filePath := filepath.Join(fileRoot, relativePath)
	return filePath
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func main() {
	url := "https://gopl-zh.github.io/index.html"
	var worklist []string
	worklist = append(worklist, url)
	breadthFirst(crawl, worklist)
}
