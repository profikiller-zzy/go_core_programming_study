package main

import (
	"fmt"
	"reflect"
)

// 专门演示反射，假设我们已经提前知道了b的具体类型是int
func reflectTest01(b interface{}) {
	//1. 先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)
	//2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	n2 := 2 + rVal.Int()
	fmt.Println("n2=", n2)
	fmt.Printf("rVal=%v rVal type=%T\n", rVal, rVal)
	//下面我们将 rVal 转成 interface{}
	iV := rVal.Interface()
	//将 interface{} 通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Println("num2=", num2)
}

// 专门演示反射[对结构体的反射]
func reflectTest02(b interface{}) {
	//通过反射获取的传入的变量的 type , kind, 值
	//1. 先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)
	//2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	//下面我们将 rVal 转成 interface{}
	iV := rVal.Interface()
	fmt.Printf("iv=%v iv type=%T \n", iV, iV)
	//将 interface{} 通过断言转成需要的类型
	//这里，我们就简单使用了一带检测的类型断言. //同学们可以使用 swtich 的断言形式来做的更加的灵活
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}
}

type Student struct {
	Name string
	Age  int
}

type Monster struct {
	Name string
	Age  int
}

// 注意这里传来的i是指针，并不是一般数据类型
func setIntegerValue(i interface{}) {
	rVal := reflect.ValueOf(i)
	switch rVal.Elem().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		rVal.Elem().SetInt(64)
	default:
		panic("不是整数类型指针")
	}
}

func main() {
	//请编写一个案例，
	//演示对(基本数据类型、interface{}、reflect.Value)进行反射的基本操作
	//1. 先定义一个 int
	// var num int = 100
	// reflectTest01(num)
	//2. 定义一个 Student 的实例
	stu := Student{
		Name: "tom", Age: 20}
	reflectTest02(stu)

	num := 12
	var numPtr *int = &num
	setIntegerValue(numPtr)
	fmt.Println(*numPtr)
	fmt.Println(num)
}
