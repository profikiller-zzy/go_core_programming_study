package main

import (
	"fmt"
	"go_core_programming/code/chapter_8/account/account"
)

func main() {
	//创建一个 account 变量
	account := account.NewAccount("jzh11111", "123456", 40)
	if account != nil {
		fmt.Println("创建成功=", *account)
	} else {
		fmt.Println("创建失败")
	}
}
