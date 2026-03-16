package main

func countBits(n int) []int {
	result := make([]int, n+1)
	for i := 0; i <= n; i++ {
		count, num := 0, i
		for num != 0 {
			if num&1 == 1 {
				count++
			}
			num /= 2
		}
		result[i] = count
	}
	return result
}
