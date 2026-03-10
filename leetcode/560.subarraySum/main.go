package main

func subarraySum(nums []int, k int) int {
	var count, preSum int
	// preHash 记录某个前缀和的个数，key为该前缀和的值，value是该前缀和的个数
	preHash := make(map[int]int)
	preHash[0] = 1
	for _, num := range nums {
		preSum += num
		if _, ok := preHash[preSum-k]; ok {
			count += preHash[preSum-k]
		}
		preHash[preSum]++
	}
	return count
}

func main() {

}
