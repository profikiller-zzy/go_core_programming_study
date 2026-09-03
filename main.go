package main

import "fmt"

func RemoveDuplicates(s string) string {
	sBytes := []byte(s)
	removeIndex := make(map[int]struct{})
	var stack []int
	for i := 0; i < len(sBytes); i++ {
		ch := sBytes[i]
		switch ch {
		case '(':
			stack = append(stack, i)
		case ')':
			if len(stack) == 0 {
				removeIndex[i] = struct{}{}
			} else {
				stack = stack[:len(stack)-1]
			}
		default:
		}
	}
	for _, i := range stack {
		removeIndex[i] = struct{}{}
	}
	var result []byte
	for i := 0; i < len(sBytes); i++ {
		if _, ok := removeIndex[i]; !ok {
			result = append(result, sBytes[i])
		}
	}
	return string(result)
}

func main() {
	s := "(a)())()))(("
	fmt.Println(RemoveDuplicates(s))
}
