package main

import (
	"bufio"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// 练习 7.17： 扩展xmlselect程序以便让元素不仅可以通过名称选择，也可以通过它们CSS风格的属性进行选择。例如一个像这样
//
//
// <div id="page" class="wide">
// 的元素可以通过匹配id或者class，同时还有它的名称来进行选择。

// attr 定义结构体用于存储需要指定的CSS风格属性
type attr struct {
	Name  string
	Value string
}

func main() {
	// 创建一个 XML 解码器，读取标准输入
	//htmlFile, err := os.OpenFile("C:\\Users\\hasee\\GolandProjects\\goStudy\\Answer\\ch7\\xml\\html.txt", os.O_RDONLY, 0666)
	//if err != nil {
	//	log.Fatal(err)
	//}
	dec := xml.NewDecoder(os.Stdin)
	var stack []string // 存放元素名称和属性的栈

	attrList, err := readAttr()
	if err != nil {
		fmt.Println(err.Error())
		attrList = nil
	}
	var currentAttrStack [][]xml.Attr // 存储最近的xml.StartElement标签中的属性信息

	for {
		// 获取下一个 XML 令牌
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}

		// 通过类型断言处理不同类型的 XML 令牌
		switch tok := tok.(type) {
		case xml.StartElement:
			// 构造包含元素名称和属性的字符串，并压入栈中
			stack = append(stack, tok.Name.Local)
			currentAttrStack = append(currentAttrStack, tok.Attr)
		case xml.EndElement:
			// 从栈中弹出元素名称和属性
			stack = stack[:len(stack)-1]
			currentAttrStack = currentAttrStack[:len(currentAttrStack)-1]
		case xml.CharData:
			// 如果栈包含命令行参数中的所有选择器，则打印文本内容
			if containsAll(stack, os.Args[1:]) && isContainAttr(currentAttrStack[len(currentAttrStack)-1], attrList) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if matches(x[0], y[0]) {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

// matches reports whether an element string matches a selector.
func matches(element, selector string) bool {
	// 检查选择器是否在元素字符串中
	return strings.Contains(element, selector)
}

// readAttr 读取需要指定的CSS风格属性
func readAttr() ([]attr, error) {
	//fmt.Println("请输入要匹配的XML属性，格式为 id=\"page\"，如果有多个输入请将它们用空格分开：")
	var attrList []attr

	// 读取用户输入
	file, err := os.OpenFile("C:\\Users\\hasee\\GolandProjects\\goStudy\\Answer\\ch7\\xml\\attr.txt", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("打开文件失败", err)
		return nil, err
	}
	reader := bufio.NewReader(file)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取输入时出错:", err)
		return nil, err
	}

	// 去掉输入字符串的换行符
	input = strings.TrimSpace(input)
	attrStrList := strings.Split(input, " ")
	for _, str := range attrStrList {
		if str == "" { // 用户输入了空字符串
			return nil, nil
		}
		parts := strings.SplitN(str, "=", 2)
		if len(parts) != 2 {
			return nil, errors.New("非法输入格式，请确保你的输入格式为 id=\"page\"")
		}
		// 去掉键和值的引号和空格
		key := strings.TrimSpace(parts[0])
		value := strings.Trim(strings.TrimSpace(parts[1]), `"`)
		attrList = append(attrList, attr{
			Name:  key,
			Value: value,
		})

		fmt.Println(key)
		fmt.Println(value)
	}
	return attrList, nil
}

// isContainAttr 用来检测attr 是否包含attr1中表示的属性对
func isContainAttr(attr []xml.Attr, attr1 []attr) bool {
	if len(attr1) == 0 {
		return true
	}
	attrMap := make(map[string]string, 0)
	for _, attrx := range attr { // 构建map
		attrMap[attrx.Name.Local] = attrx.Value
	}
	for _, attry := range attr1 {
		if value, ok := attrMap[attry.Name]; !ok || value != attry.Value {
			return false
		}
	}
	return true
}
