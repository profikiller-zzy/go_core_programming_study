package main

import "fmt"

func main() {
	var f float64
	f = 64.0
	var i interface{}
	i = f
	value, ok := i.(float32)
	if ok { // 断言成功
		fmt.Printf("接口 i 的动态类型为%T, 动态值为%v", value, value)
	} else {
		fmt.Println("断言失败，接口 i 的动态类型不是 float32")
	}
}
