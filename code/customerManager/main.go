package main

import (
	"go_core_programming/code/customerManager/customerView"
	"go_core_programming/code/customerManager/service"
)

func main() {
	//在 main 函数中，创建一个 customerView,并运行显示主菜单..
	customerView := customerView.NewCustomerView("", true)
	//这里完成对 customerView 结构体的 customerService 字段的初始化
	customerView.SetCustomerService(service.NewCustomerService())
	//显示主菜单..
	customerView.MainMenu()
}
