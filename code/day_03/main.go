package main

import (
	"fmt"
	"strconv"
)

func main() {
	numStr := "1234567"
	num, err := strconv.ParseInt(numStr, 10, 16)
	fmt.Println(num, err)

	floatStr := "1103bba"
	floatNum, err := strconv.ParseFloat(floatStr, 32)
	fmt.Println(floatNum, err)

	var i int = 10
	fmt.Printf("%v\n", &i)
}
