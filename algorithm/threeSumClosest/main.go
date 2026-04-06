package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	// 先将nums排序
	sort.Ints(nums)
	abs := func(n int) int { // 匿名辅助函数，返回一个数的绝对值
		if n < 0 {
			return -n
		}
		return n
	}
	// 初始化最小距离和对应的三元组之和
	minDis := abs(nums[0] + nums[1] + nums[2] - target)
	minDisSum := nums[0] + nums[1] + nums[2]

	for left := 0; left < len(nums)-2; left++ {
		if left > 0 && nums[left] == nums[left-1] { // 跳过重复的
			continue
		}
		mid, right := left+1, len(nums)-1
		for mid < right {
			sum := nums[left] + nums[mid] + nums[right]
			diff := sum - target
			if abs(diff) < minDis { // 找到更接近target的三元组
				minDisSum = sum
				minDis = abs(diff)
			}

			if diff == 0 {
				return sum
			} else if diff > 0 { // 这种情况可能是右边的数太大了
				right--
			} else {
				mid++
			}
		}
	}
	return minDisSum
}

func main() {
	var nums []int
	nums = []int{4, 0, 5, -5, 3, 3, 0, -4, -5}
	target := -2
	fmt.Println(threeSumClosest(nums, target))
}
