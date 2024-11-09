package main

import (
	"fmt"
	"strings"
)

func main() {
	index := strings.LastIndex("go golang", "go")
	fmt.Println(index)

	str := strings.Replace("go go golang", "go", "背景", -1)
	fmt.Println(str)

	strArr := strings.Split("I love you", " ")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("strArr[%d]: %s\n", i, strArr[i])
	}
	fmt.Printf("strArr： %v", strArr)

	fmt.Println(strings.TrimSpace(" tn a lone gopher ntrn     "))

	fmt.Println(strings.Trim("! hello    ! !!! ", "! "))

	strings.HasPrefix("ftp://192.168.10.1", "ftp") // true
}
