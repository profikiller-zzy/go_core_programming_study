package main

import (
	"fmt"
	"reflect"
)

// Monster 定义了一个 Monster 结构体
type Monster struct {
	Name  string  `json:"name"`
	Age   int     `json:"monster_age"`
	Score float32 `json:"成绩"`
	Sex   string
}

// GetSum 方法，返回两个数的和
func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

// Set 方法， 接收四个值，给 s 赋值
func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

// Print 方法，显示 s 的值
func (s Monster) Print() {
	fmt.Println("---start~----")
	fmt.Println(s)
	fmt.Println("---end~----")
}

func TestStruct(a interface{}) {
	//获取 reflect.Type 类型
	typ := reflect.TypeOf(a)
	//获取 reflect.Value 类型
	val := reflect.ValueOf(a)
	//获取到 a 对应的类别
	kd := val.Kind()
	//如果传入的不是 struct，就退出
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	//获取到该结构体有几个字段
	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num) //4
	//变量结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: 值为=%v\n", i, val.Field(i))
		//获取到 struct 标签, 注意需要通过 reflect.Type 来获取 tag 标签的值
		tagVal := typ.Field(i).Tag.Get("json")
		//如果该字段于 tag 标签就显示，否则就不显示
		if tagVal != "" {
			fmt.Printf("Field %d: tag 为=%v\n", i, tagVal)
		}
	}
	//获取到该结构体有多少个方法
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)
	//var params []reflect.Value
	//方法的排序默认是按照 函数名的排序（ASCII 码）
	val.Method(1).Call(nil) //获取到第二个方法并调用，这里的方法是Print
	// 								   排序的顺序是0.GetSum 1.Print 2.Set
	//调用结构体的第 1 个方法 0.GetSum
	var params []reflect.Value //声明了 []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params) //传入的参数是 []reflect.Value, 返回[]reflect.Value
	fmt.Println("res=", res[0].Int()) //返回结果, 返回的结果是 []reflect.Value*/
}

func main() {
	//创建了一个 Monster 实例
	var a Monster = Monster{
		Name: "黄鼠狼精", Age: 400, Score: 30.8,
	}
	//将 Monster 实例传递给 TestStruct 函数
	TestStruct(a)
}
