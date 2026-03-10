package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	// 检查是否有命令行参数
	if len(os.Args) < 2 {
		fmt.Println("用法: go run main.go <要编码的文本>")
		fmt.Println("示例: go run main.go \"Hello World\"")
		os.Exit(1)
	}

	// 获取输入文本（支持带空格的参数，需将多个参数合并或只取第一个）
	// 这里我们将所有参数用空格连接，以支持 "Hello World" 这种输入
	inputText := ""
	for i, arg := range os.Args[1:] {
		if i > 0 {
			inputText += " "
		}
		inputText += arg
	}

	// 将字符串转换为字节切片
	inputBytes := []byte(inputText)

	// 执行 Base64 编码
	encodedStr := base64.StdEncoding.EncodeToString(inputBytes)

	// 输出结果
	fmt.Println(encodedStr)
}
