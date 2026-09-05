package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func solution(budget, capacity, carCost, areaCost int, height []int) int {
	var total int
	// 直接计算可以将多少土填埋进去
	for budget > 0 {
		if budget < carCost {
			break
		}
		budget -= carCost
		minArea := min(capacity, budget/areaCost)
		budget -= minArea * areaCost
		total += minArea
	}
	sort.Slice(height, func(i, j int) bool {
		return height[i] < height[j]
	})
	minHeight := height[0]
	var index int
	for {
		for index < len(height) && height[index] <= minHeight {
			index++
		}
		if total < index {
			break
		}
		total -= index
		minHeight++
	}
	return minHeight
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var budget, capacity, carCost, areaCost, n int
	fmt.Fscan(in, &budget, &capacity, &carCost, &areaCost, &n)
	height := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &height[i])
	}
	fmt.Println(solution(budget, capacity, carCost, areaCost, height))
}
