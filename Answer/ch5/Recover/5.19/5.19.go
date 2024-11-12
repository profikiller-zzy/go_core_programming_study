package main

import "fmt"

// 练习5.19： 使用panic和recover编写一个不包含return语句但能返回一个非零值的函数。

// 具体做法是在函数中使用`panic`来引发一个异常，并在延迟函数中使用`recover`
// 捕获这个异常，并返回非零值

func getValue() (value int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			value = 1
		}
	}()

	panic(42) // 引发 panic，传递一个非零值
}

func main() {
	value := getValue()
	fmt.Println("Returned value:", value)
}
