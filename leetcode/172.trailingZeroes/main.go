package main

func trailingZeroes(n int) int {
	var ans int
	for index := 5; index <= n; index += 5 {
		for x := index; x%5 == 0; x /= 5 {
			ans++
		}
	}
	return ans
}
