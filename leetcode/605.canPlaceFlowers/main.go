package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)
	var canFlowerbed int
	for i, isFlowerbed := range flowerbed {
		if isFlowerbed == 1 {
			continue
		}
		prev := i > 0 && flowerbed[i-1] == 1
		next := i < length-1 && flowerbed[i+1] == 1
		if !prev && !next {
			canFlowerbed++
			flowerbed[i] = 1
		}
	}
	return canFlowerbed >= n
}

func main() {
	flowerbed := []int{1, 0, 0, 0, 0, 1}
	fmt.Println(canPlaceFlowers(flowerbed, 2))
}
