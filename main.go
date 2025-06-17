package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	str := "测试"
	bytes, err := json.Marshal(str)
	if err != nil {
		fmt.Println("序列化失败:", err)
		return
	}
	fmt.Println(string(bytes))
}
