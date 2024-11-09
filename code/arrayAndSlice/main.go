package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5} // 长度会自动推断为 5
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	intArr := [5]int{1, 22, 33, 44, 55}
	intSlice := intArr[2:4]
	for i := 0; i < 5; i++ {
		fmt.Printf("intArr[%d]地址:%v\n", i, &intArr[i])
	}
	for i := 0; i < len(intSlice); i++ {
		fmt.Printf("intSlice[%d]地址:%v\n", i, &intSlice[i])
	}
	fmt.Println(cap(intSlice))
	fmt.Println(len(intSlice))
}
