package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Hero 结构体
type Hero struct {
	Name string
	Age  int
}

// HeroSlice Hero 结构体切片类型
type HeroSlice []Hero

// 实现 sort.Interface 接口

// Len sort.Interface 接口的第一个方法
func (hs HeroSlice) Len() int {
	return len(hs)
}

// Less 方法就是决定你使用什么标准进行排序
// 1. 按 Hero 的年龄从小到大排序!!
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
	//修改成对 Name 排序
	//return hs[i].Name < hs[j].Name
}

func (hs HeroSlice) Swap(i, j int) {
	//交换
	// temp := hs[i]
	// hs[i] = hs[j]
	// hs[j] = temp
	//下面的一句话等价于三句话
	hs[i], hs[j] = hs[j], hs[i]
}

// Student 结构体
type Student struct {
	Name  string
	Age   int
	Score float64
}

// 将 Student 的切片，按 Score 从大到小排序!!
func main() {
	//先定义一个数组/切片
	var intSlice = []int{0, -1, 10, 7, 90}
	//要求对 intSlice 切片进行排序
	//1. 冒泡排序... //2. 也可以使用系统提供的方法
	sort.Ints(intSlice)
	fmt.Println(intSlice)
	//请大家对结构体切片进行排序
	//1. 冒泡排序... //2. 也可以使用系统提供的方法
	//测试看看我们是否可以对结构体切片进行排序
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄|%d", rand.Intn(100)), Age: rand.Intn(100)}
		//将 hero append 到 heroes 切片
		heroes = append(heroes, hero)
	}
	//看看排序前的顺序
	for _, v := range heroes {
		fmt.Println(v)
	}
	//调用 sort.Sort
	sort.Sort(heroes)
	fmt.Println("-----------排序后------------")
	//看看排序后的顺序
	for _, v := range heroes {
		fmt.Println(v)
	}
}
