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

	intSlice[0] = 2
	for i := 0; i < 5; i++ {
		fmt.Printf("%d\t", intArr[i])
	}
	fmt.Println()
	for i := 0; i < len(intSlice); i++ {
		fmt.Printf("%d\t", intSlice[i])
	}
	fmt.Println()

	var floatArray [5]float64
	for i, _ := range floatArray {
		floatArray[i] = float64(i)
	}
	var floatSlice1 []float64 = floatArray[0:5]
	//var floatSlice2 []float64 = make([]float64, 5, 10)
	fmt.Printf("len floatSlice1:%d\n", len(floatSlice1))
	fmt.Printf("cap floatSlice1:%d\n", cap(floatSlice1))
	fmt.Printf("len floatArray:%d\n", len(floatArray))
	floatSlice1 = append(floatSlice1, 5, 6)
	fmt.Printf("after append: len floatSlice1:%d\n", len(floatSlice1))
	fmt.Printf("after append: cap floatSlice1:%d\n", cap(floatSlice1))
	fmt.Printf("after append: len floatArray:%d\n", len(floatArray))
}
