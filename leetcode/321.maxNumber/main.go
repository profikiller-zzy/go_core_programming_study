package main

import "fmt"

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	m, n := len(nums1), len(nums2)
	var result []int

	// 枚举所有可能的k1取值
	for k1 := max(0, k-n); k1 <= min(k, m); k1++ {
		k2 := k - k1

		// 获取两个数组的最大子序列
		sub1 := maxSubsequence(nums1, k1)
		sub2 := maxSubsequence(nums2, k2)

		// 合并两个子序列
		merged := merge(sub1, sub2)

		// 选择字典序最大的结果
		if isLarger(merged, result) {
			result = merged
		}
	}

	return result
}

func maxSubsequence(nums []int, k int) []int {
	if k == 0 {
		return []int{}
	}

	stack := make([]int, 0, k)
	toRemove := len(nums) - k // 需要删除的元素额度

	for _, num := range nums {
		// 当栈不为空且当前元素大于栈顶元素，且还有删除额度
		for len(stack) > 0 && num > stack[len(stack)-1] && toRemove > 0 {
			stack = stack[:len(stack)-1]
			toRemove--
		}
		stack = append(stack, num)
	}

	// 如果还有移除额度，从末尾移除
	if toRemove > 0 {
		stack = stack[:len(stack)-toRemove]
	}

	return stack
}

// 合并两个子序列形成最大数组
func merge(nums1 []int, nums2 []int) []int {
	result := make([]int, 0, len(nums1)+len(nums2))
	i, j := 0, 0

	for i < len(nums1) && j < len(nums2) {
		// 比较字典序，选择较大的元素
		if compare(nums1, i, nums2, j) {
			result = append(result, nums1[i])
			i++
		} else {
			result = append(result, nums2[j])
			j++
		}
	}

	// 添加剩余元素
	result = append(result, nums1[i:]...)
	result = append(result, nums2[j:]...)

	return result
}

// compare 比较两个子序列分别从i, j元素开始的字典序
func compare(nums1 []int, i int, nums2 []int, j int) bool {
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] != nums2[j] {
			return nums1[i] > nums2[j]
		}
		i++
		j++
	}
	// 如果一个序列是另一个序列的前缀，选择较长的序列
	return i < len(nums1)
}

// isLarger 判断新结果是否比当前结果大
func isLarger(new, old []int) bool {
	if len(old) == 0 {
		return true
	}
	for i := 0; i < len(new); i++ {
		if new[i] != old[i] {
			return new[i] > old[i]
		}
	}
	return false
}

func main() {
	nums1 := []int{9, 1, 2, 5, 8, 3}
	nums2 := []int{3, 4, 6, 5}
	k := 5
	result := maxNumber(nums1, nums2, k)
	fmt.Println(result)
}
