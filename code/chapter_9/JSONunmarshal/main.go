package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string
	Age      int
	Birthday string //.... Sal float64
	Skill    string
}

// 演示将 json 字符串，反序列化成 struct
func unmarshalStruct() {
	//说明 str 在项目开发中，是通过网络传输获取到.. 或者是读取文件获取到
	str := "{\"Name\":\"牛魔王\",\"Age\":500,\"Birthday\":\"2011-11-11\",\"Sal\":8000,\"Skill\":\"牛魔拳\"}"
	//定义一个 Monster 实例
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 monster=%v monster.Name=%v \n", monster, monster.Name)
}

// 演示将 json 字符串，反序列化成 map
func unmarshalMap() {
	str := "{\"address\":\"洪崖洞\",\"age\":30,\"name\":\"红孩儿\"}"
	//定义一个 map
	var a map[string]interface{}
	//反序列化
	//注意：反序列化 map,不需要 make,因为 make 操作被封装到 Unmarshal 函数
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 a=%v\n", a)
}

// 演示将 json 字符串，反序列化成切片
func unmarshalSlice() {
	str := "[{\"address\":\"北京\",\"age\":\"7\",\"name\":\"jack\"}," +
		"{\"address\":[\"墨西哥\",\"夏威夷\"],\"age\":\"20\",\"name\":\"tom\"}]"
	//定义一个 slice
	var slice []map[string]interface{}
	//反序列化，不需要 make,因为 make 操作被封装到 Unmarshal 函数
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 slice=%v\n", slice)
}

func main() {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}
