package main

import (
	"fmt"
	"strings"
)

/*
闭包的最佳实践

请编写一个程序，具体要求如下：

编写一个函数 makeSuffix(suffix string)，可以接收一个文件后缀名（比如 .jpg），并返回一个闭包。
调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀（比如 .jpg），则返回 文件名.jpg，如果已经有 .jpg 后缀，则返回原文件名。
要求使用闭包的方式完成。
使用 strings.HasSuffix
*/
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		// 判断文件名是否以指定后缀结尾
		if !strings.HasSuffix(name, suffix) {
			// 如果没有指定后缀，添加后缀
			return name + suffix
		}
		// 有后缀，直接返回源文件名
		return name
	}
}

func main() {
	// 定义一个函数闭包
	addJpgSuffix := makeSuffix(".jpg")
	fmt.Println(addJpgSuffix("file"))      // 输出 file.jpg
	fmt.Println(addJpgSuffix("image.jpg")) // 输出 image.jpg
}
