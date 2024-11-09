package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 打开二进制文件
	binaryFilePath := "C:\\Users\\hasee\\GolandProjects\\go_core_programming\\code\\db_description\\db_description" // 二进制文件路径
	data, err := ioutil.ReadFile(binaryFilePath)
	if err != nil {
		fmt.Println("Error reading binary file:", err)
		return
	}

	// 创建或打开TXT文件
	txtFilePath := "output.txt" // 输出TXT文件路径
	txtFile, err := os.Create(txtFilePath)
	if err != nil {
		fmt.Println("Error creating TXT file:", err)
		return
	}
	defer txtFile.Close()

	// 将字节写入TXT文件
	_, err = txtFile.Write(data)
	if err != nil {
		fmt.Println("Error writing to TXT file:", err)
		return
	}

	fmt.Println("Successfully written binary data to", txtFilePath)
}
