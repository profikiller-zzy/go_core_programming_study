package main

func candy(ratings []int) int {
	candyNum := make([]int, len(ratings))
	// 首先初始化
	for index := 0; index < len(candyNum); index++ {
		candyNum[index] = 1
	}
	for index := 0; index < len(candyNum); index++ {
		theWeak := make([]int, 0)
		if index-1 >= 0 && ratings[index] > ratings[index-1] {
			theWeak = append(theWeak, candyNum[index-1])
		}
		if index+1 < len(candyNum) && ratings[index] > ratings[index+1] {
			theWeak = append(theWeak, candyNum[index+1])
		}
		if len(theWeak) == 0 {
			continue
		} else if len(theWeak) == 1 {
			candyNum[index] = theWeak[0] + 1
		} else {
			candyNum[index] = max(theWeak[0], theWeak[1]) + 1
		}
	}
	var total int
	for index := 0; index < len(candyNum); index++ {
		total += candyNum[index]
	}
	return total
}

func main() {

}
