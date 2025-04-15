package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())
	s := make([]int, 10, 20)
	for i := 0; i < 10; i++ {
		s[i] = i
	}
	s1 := append(s[0:5], s[6:7]...)
	fmt.Printf("slice: %v, len: %v, cap: %v\n", s1, len(s1), cap(s1))
	s1[0] = 9
	fmt.Println(s)
	testSlice()
}

func testSlice() {
	s := make([]int, 512)
	s = append(s, 1)
	fmt.Printf("len: %v, cap: %v\n", len(s), cap(s))
}
