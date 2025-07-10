package main

import "fmt"

func hammingWeight1(n int) int {
	if n == 0 {
		return 0
	}
	var res int
	for n != 0 {
		num2power := return2power(n)
		n -= num2power
		res++
	}
	return res
}

func return2power(num int) int {
	res := 1
	if res == num {
		return res
	}
	for num >= res*2 {
		res *= 2
	}
	return res
}

func hammingWeight2(n int) int {
	var res int
	for i := 0; i < 32; i++ {
		if n&(1<<i) != 0 {
			res++
		}
	}
	return res
}

func hammingWeight(n int) int {
	var res int
	for ; n != 0; n &= n - 1 { // n &= n - 1 的作用是将 n 的最低位的 1 变为 0，其他位不变
		res++
	}
	return res
}

func main() {
	n := 128
	fmt.Println(hammingWeight(n))
}
