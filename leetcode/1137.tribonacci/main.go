package main

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	t := make([]int, n+1)
	t[0] = 0
	t[1], t[2] = 1, 1
	for i := 3; i <= n; i++ {
		t[i] = t[i-3] + t[i-2] + t[i-1]
	}
	return t[n]
}
