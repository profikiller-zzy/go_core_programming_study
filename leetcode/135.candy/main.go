package main

func candy(ratings []int) int {
	n := len(ratings)
	candies := make([]int, n)

	// 每人先分 1 个糖果
	for i := 0; i < n; i++ {
		candies[i] = 1
	}

	// 从左往右：如果右边比左边评分高，右边糖果 = 左边 + 1
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// 从右往左：如果左边比右边评分高，左边糖果 = max(左边当前, 右边 + 1)
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}

	// 累加所有糖果
	total := 0
	for _, c := range candies {
		total += c
	}
	return total
}

func main() {

}
