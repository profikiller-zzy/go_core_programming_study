package main

func findDuplicate(nums []int) int {
	// 索引法 index 0 1 ... n-1
	//       value 1 2 ... n
	n := len(nums) - 1
	if n == 1 {
		return nums[0]
	}
	for index := 0; index < n; index++ {
		for nums[index] != index+1 {
			if nums[index] == nums[nums[index]-1] {
				return nums[index]
			}
			// 交换 nums[index] 和 nums[nums[index]-1]
			nums[nums[index]-1], nums[index] = nums[index], nums[nums[index]-1]
		}
	}
	return nums[n]
}

func main() {

}
