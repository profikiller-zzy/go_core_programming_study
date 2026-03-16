package main

import "fmt"

func minFlips(a int, b int, c int) int {
	var count int
	for a != 0 || b != 0 || c != 0 {
		abit, bbit, cbit := a&1, b&1, c&1
		if abit|bbit != cbit {
			if cbit == 0 { // 那么 abit or bbit 的结果一定是1
				count += abit + bbit
			} else {
				count++
			}
		}
		a /= 2
		b /= 2
		c /= 2
	}
	return count
}

func main() {
	fmt.Println(minFlips(2, 6, 5))
}
