package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// 实现 String 方法
func (p Person) String() string {
	return fmt.Sprintf("Person{Name: %s, Age: %d}", p.Name, p.Age)
}

func main() {
	p := Person{Name: "Alice", Age: 30}

	// fmt.Println 会调用 String 方法
	fmt.Println(p) // 输出：Person{Name: Alice, Age: 30}
}
