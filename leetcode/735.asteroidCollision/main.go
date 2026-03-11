package main

import "fmt"

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0, len(asteroids))
	for _, a := range asteroids {
		survived := true
		for len(stack) > 0 && stack[len(stack)-1] > 0 && a < 0 && survived {
			m1, m2 := stack[len(stack)-1], -a // m1, m2分别是栈顶行星和当前行星的质量
			if m1 > m2 {
				survived = false
			} else if m1 < m2 {
				stack = stack[:len(stack)-1]
			} else {
				stack = stack[:len(stack)-1]
				survived = false
			}
		}
		if survived {
			stack = append(stack, a)
		}
	}
	return stack
}

func main() {
	asteroids := []int{3, 5, -6, 2, -1, 4}
	fmt.Println(asteroidCollision(asteroids))
}
