package main

func findDifference(nums1 []int, nums2 []int) [][]int {
	numsSet1, numsSet2 := make(map[int]interface{}), make(map[int]interface{})
	for _, num := range nums1 {
		numsSet1[num] = nil
	}
	for _, num := range nums2 {
		numsSet2[num] = nil
	}

	result := make([][]int, 2)
	for num := range numsSet1 {
		if _, ok := numsSet2[num]; !ok {
			result[0] = append(result[0], num)
		}
	}
	for num := range numsSet2 {
		if _, ok := numsSet1[num]; !ok {
			result[1] = append(result[1], num)
		}
	}
	return result
}
