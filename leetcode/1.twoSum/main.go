package main

func towSum(nums []int, target int) []int {
	res := []int{}
	for i := 0; i < len(nums)-1; i++ { // 加数1
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res = []int{i, j}
			}
		}
	}
	return res
}

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
