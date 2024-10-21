package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num int8 = -128
	var ch rune = 1
	var num1 int
	fmt.Printf("%x\n", num)
	fmt.Printf("Size of rune: %d bytes\n", unsafe.Sizeof(ch))
	fmt.Printf("Size of int: %d bytes\n", unsafe.Sizeof(num1))

	var x int64
	fmt.Printf("x 的类型为 %T\nx 所占的字节数为 %d\n", x, unsafe.Sizeof(x))

	var y = 100
	fmt.Printf("y 的类型为 %T\n", y)

	var z = 1.01
	fmt.Printf("z 的类型为 %T\n", z)
}
