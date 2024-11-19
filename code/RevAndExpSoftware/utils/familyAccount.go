package utils

import (
	"fmt"
)

type FamilyAccount struct {
	//声明必须的字段. //声明一个字段，保存接收用户输入的选项
	key string
	//声明一个字段，控制是否退出 for
	loop bool
	//定义账户的余额 []
	balance float64
	//每次收支的金额
	money float64
	//每次收支的说明
	note string
	//定义个字段，记录是否有收支的行为
	flag bool
	//收支的详情使用字符串来记录
	//当有收支时，只需要对 details 进行拼接处理即可
	details string
}

// NewFamilyAccount 编写要给工厂模式的构造方法，返回一个*FamilyAccount 实例
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key: "", loop: true, balance: 10000.0, money: 0.0, note: "", flag: false, details: "收支\t 账户金额\t 收支金额\t 说 明"}
}

// showDetails 将显示明细写成一个方法
func (fa *FamilyAccount) showDetails() {
	fmt.Println("-----------------当前收支明细记录-----------------")
	if fa.flag {
		fmt.Println(fa.details)
	} else {
		fmt.Println("当前没有收支明细... 来一笔吧!")
	}
}

// 将登记收入写成一个方法，和*FamilyAccount 绑定
func (fa *FamilyAccount) income() {
	fmt.Println("本次收入金额:")
	fmt.Scanln(&fa.money)
	fa.balance += fa.money // 修改账户余额
	fmt.Println("本次收入说明:")
	fmt.Scanln(&fa.note)
	//将这个收入情况，拼接到 details 变量
	//收入 11000 1000 有人发红包
	fa.details += fmt.Sprintf("\n 收入\t%v\t%v\t%v", fa.balance, fa.money, fa.note)
	fa.flag = true
}

// 将登记支出写成一个方法，和*FamilyAccount 绑定
func (fa *FamilyAccount) pay() {
	fmt.Println("本次支出金额:")
	fmt.Scanln(&fa.money)
	//这里需要做一个必要的判断
	if fa.money > fa.balance {
		fmt.Println("余额的金额不足")
		//break
	}
	fa.balance -= fa.money
	fmt.Println("本次支出说明:")
	fmt.Scanln(&fa.note)
	fa.details += fmt.Sprintf("\n 支出\t%v\t%v\t%v", fa.balance, fa.money, fa.note)
	fa.flag = true
}

// 将退出系统写成一个方法,和*FamilyAccount 绑定
func (fa *FamilyAccount) exit() {
	fmt.Println("你确定要退出吗? y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入 y/n")
	}
	if choice == "y" {
		fa.loop = false
	}
}

// 给该结构体绑定相应的方法
// 显示主菜单
func (fa *FamilyAccount) MainMenu() {
	for {
		fmt.Println("\n-----------------家庭收支记账软件-----------------")
		fmt.Println(" 1 收支明细")
		fmt.Println(" 2 登记收入")
		fmt.Println(" 3 登记支出")
		fmt.Println(" 4 退出软件")
		fmt.Print("请选择(1-4)：")
		fmt.Scanln(&fa.key)
		switch fa.key {
		case "1":
			fa.showDetails()
		case "2":
			fa.income()
		case "3":
			fa.pay()
		case "4":
			fa.exit()
		default:
			fmt.Println("请输入正确的选项..")
		}
		if !fa.loop {
			break
		}
	}
}
