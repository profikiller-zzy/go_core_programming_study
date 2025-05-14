package main

import "fmt"

func nextPermutation(nums []int) {
	reverse := func(start int, end int) {
		for i := start; i <= (start+end)/2; i++ {
			// 计算对称位置的索引
			j := end - (i - start)
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	for i := len(nums) - 2; i >= 0; i-- {
		// 从后往前找到第一个升序的位置，这说明找到了需要调整的位置，即[i:len(nums)-1]
		if nums[i] < nums[i+1] {
			for j := len(nums) - 1; j > i; j-- {
				// 从后往前找到第一个比nums[i]大的位置，交换两个位置的值
				if nums[j] > nums[i] {
					nums[i], nums[j] = nums[j], nums[i]
					// 此时 i+1 到结尾一定是降序的；交换过后也一定是降序的，反转让 i+1 到结尾变为升序，变成最小的排列，这样保证增量是最小的，也就找到了下一个排列
					reverse(i+1, len(nums)-1)
					return
				}
			}
		}
	}
	// 如果没有找到升序的位置，说明整个数组是降序的，直接反转数组即可
	reverse(0, len(nums)-1)
	return
}

func main() {
	nums := []int{3, 2, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}
