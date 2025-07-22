package main

import (
	"fmt"
)

func main() {
	x := 42 // 示例数字
	binaryString := fmt.Sprintf("%b", x)
	fmt.Println("The binary representation of", x, "is:", binaryString)
}
