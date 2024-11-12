package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := Person{Name: "Alice", Age: 30}

	// 序列化为 JSON
	jsonData, _ := json.Marshal(p)
	fmt.Println(string(jsonData)) // 输出：{"name":"Alice","age":30}

	// 反序列化 JSON
	var p2 Person
	json.Unmarshal(jsonData, &p2)
	fmt.Println(p2) // 输出：{Alice 30}
}
