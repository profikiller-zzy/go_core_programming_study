package main

func maxOperations(nums []int, k int) int {
	numCount := make(map[int]int)
	count := 0

	for _, num := range nums {
		complement := k - num
		if numCount[complement] > 0 {
			numCount[complement]--
			count++
		} else {
			numCount[num]++
		}
	}

	return count
}
