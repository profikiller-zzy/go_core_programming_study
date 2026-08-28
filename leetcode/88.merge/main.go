package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	result := make([]int, 0, m+n)
	var cur1, cur2 int
	for {
		if cur1 == m {
			result = append(result, nums2[cur2:]...)
			break
		}
		if cur2 == n {
			result = append(result, nums1[cur1:]...)
			break
		}
		if nums1[cur1] < nums2[cur2] {
			result = append(result, nums1[cur1])
			cur1++
		} else {
			result = append(result, nums2[cur2])
			cur2++
		}
	}
	copy(nums1, result)
}
