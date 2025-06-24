package main

func containsNearbyDuplicate(nums []int, k int) bool {
	numHash := map[int]int{}
	for index, value := range nums {
		if index1, ok := numHash[value]; ok && index-index1 <= k {
			return true
		} else {
			numHash[value] = index
		}
	}
	return false
}
