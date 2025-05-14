package main

func singleNumber1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	numHash := make(map[int]bool)
	for _, num := range nums {
		if _, ok := numHash[num]; ok {
			delete(numHash, num)
		} else {
			numHash[num] = true
		}
	}

	var res int
	for num, _ := range numHash {
		res = num
		break
	}
	return res
}

func singleNumber(nums []int) int {
	var res int
	for _, num := range nums {
		res = res ^ num
	}
	return res
}

func main() {

}
