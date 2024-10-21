package main

import "fmt"

// go 语言指针的使用特点
func testPtr(num *int) {
	*num = 20 // * 解引用为这个指针赋值
}
func sayOK() {
	// 输出一句话
	fmt.Println("ok")
}

func sumAndSub(a int, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}

func printPersonalInformation(name string, sex string, age int, address string) {
	res := fmt.Sprintf("姓名: %s\n性别: %s\n年龄: %d\n家庭住址: %s\n", name, sex, age, address)
	fmt.Println(res)
}

func main() {
	printPersonalInformation("郭舒澜", "女", 24, "大冶一中")
}
