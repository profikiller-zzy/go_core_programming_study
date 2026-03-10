package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var maxCandyNum int
	for _, candyNum := range candies {
		if candyNum > maxCandyNum {
			maxCandyNum = candyNum
		}
	}

	result := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		result[i] = candies[i]+extraCandies >= maxCandyNum
	}
	return result
}
