package service

import "go_core_programming/code/customerManager/model"

// CustomerService 相当于customer数据库，完成对 Customer 的操作,包括增删改查
type CustomerService struct {
	customers   []model.Customer
	customerNum int
}

// NewCustomerService 编写一个方法，可以返回 *CustomerService
func NewCustomerService() *CustomerService {
	// 为了能够看到有客户在切片中，我们初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "112", "zs@sohu.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

// List 返回客户切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
}

// Add 添加客户到 customers 切片
func (this *CustomerService) Add(customer model.Customer) bool {
	// 我们确定一个分配 ID 的规则，就是添加的顺序
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

// FindById 根据 id 查找客户在切片中的下标，如果没有该客户，返回 -1
func (this *CustomerService) FindById(id int) int {
	index := -1
	// 遍历 this.customers 切片
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			// 找到
			index = i
			break
		}
	}
	return index
}

// Delete 删除客户信息的方法
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	// 如果 index == -1，说明没有这个客户
	if index == -1 {
		return false
	}
	// 从切片中删除一个元素
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}
