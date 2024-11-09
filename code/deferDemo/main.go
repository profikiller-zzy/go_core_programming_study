package main

import (
	"fmt"
)

func sum(n1 int, n2 int) int {
	// 当执行到 defer 时，暂时不执行，会将 defer 后面的语句压入到独立的栈 (defer 栈)
	// 当函数执行完毕后，再从 defer 栈按照先入后出的方式出栈，执行
	defer fmt.Println("ok1 n1=", n1) // defer
	defer fmt.Println("ok2 n2=", n2) // defer

	res := n1 + n2
	fmt.Println("ok3 res=", res)

	return res
}

func main() {
	res := sum(10, 20)
	fmt.Println("res=", res)
}
