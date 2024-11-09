package main

import "fmt"

func main() {
	test()
	// code ...
	fmt.Println("code...")
}

func test() {
	// 使用defer + recover 来处理函数中遇到的错误
	defer func() {
		err := recover()
		if err != nil { // 说明在函数执行过程中遇到了错误
			fmt.Printf("err=%v\n", err)
			// 这里可以针对错误进行各种各样的处理，或者是将错误信息发送给管理员等等
		}
	}()
	num1 := 1
	num2 := 0
	num3 := num1 / num2
	fmt.Println(num3)
}
