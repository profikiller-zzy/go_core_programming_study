package customerView

import (
	"fmt"
	"go_core_programming/code/customerManager/model"
	"go_core_programming/code/customerManager/service"
)

type customerView struct {
	//定义必要字段
	key string //接收用户输入... loop bool //表示是否循环的显示主菜单
	//增加一个字段 customerService
	customerService *service.CustomerService
	loop            bool
}

func NewCustomerView(key string, loop bool) *customerView {
	return &customerView{
		key:  key,
		loop: loop,
	}
}

func (this *customerView) SetCustomerService(customerService *service.CustomerService) {
	this.customerService = customerService
}

// 显示所有的客户信息
func (this *customerView) List() {
	//首先，获取到当前所有的客户信息(在切片中)
	customers := this.customerService.List()
	//显示
	fmt.Println("---------------------------客户列表---------------------------")
	fmt.Println("编号\t 姓名\t 性别\t 年龄\t 电话\t 邮箱")
	for i := 0; i < len(customers); i++ {
		//fmt.Println(customers[i].Id,"\t", customers[i].Name...)
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("\n-------------------------客户列表完成-------------------------\n\n")
}

// 显示主菜单
func (this *customerView) MainMenu() {
	for {
		fmt.Println("-----------------客户信息管理软件-----------------")
		fmt.Println(" 1 添 加 客 户")
		fmt.Println(" 2 修 改 客 户")
		fmt.Println(" 3 删 除 客 户")
		fmt.Println(" 4 客 户 列 表")
		fmt.Println(" 5 退 出")
		fmt.Print("请选择(1-5)：")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.Add()
		case "2":
			fmt.Println("修 改 客 户")
		case "3":
			this.Delete()
		case "4":
			this.List()
		case "5":
			this.loop = false
		default:
			fmt.Println("你的输入有误，请重新输入...")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("你退出了客户关系管理系统...")
}

func (this *customerView) Add() {
	fmt.Println("---------------------添加客户---------------------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电邮:")
	email := ""
	fmt.Scanln(&email)
	//构建一个新的 Customer 实例
	//注意: id 号，没有让用户输入，id 是唯一的，需要系统分配
	customer := model.NewCustomer2(name, gender, age, phone, email)
	//调用
	if this.customerService.Add(customer) {
		fmt.Println("---------------------添加完成---------------------")
	} else {
		fmt.Println("---------------------添加失败---------------------")
	}
}

// 得到用户的输入 id，删除该 id 对应的客户
func (this *customerView) Delete() {
	fmt.Println("--------------------- 删除客户 ---------------------")
	fmt.Println("请输入待删除客户编号(-1退出): ")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return // 放弃删除操作
	}

	fmt.Println("确认是否删除 (Y/N): ")
	// 这里同样可以加入一个循环判断，直到用户输入 y 或者 n 才退出
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		// 调用 customerService 的 Delete 方法
		if this.customerService.Delete(id) {
			fmt.Println("--------------------- 删除完成 ---------------------")
		} else {
			fmt.Println("---------------- 删除失败，输入的 id 不存在 ----------------")
		}
	}
}
