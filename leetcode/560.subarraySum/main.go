package main

// https://leetcode.cn/problems/subarray-sum-equals-k/description
// 参考leetcode题解： 前缀和加哈希表
func subarraySum(nums []int, k int) int {
	var count, preSum int
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
