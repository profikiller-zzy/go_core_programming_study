package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var getKthElement func(k int) int
	getKthElement = func(k int) int {
		cur1, cur2 := 0, 0
		for {
			if cur1 == len(nums1) {
				return nums2[cur2+k-1]
			}
			if cur2 == len(nums2) {
				return nums1[cur1+k-1]
			}
			if k == 1 {
				return min(nums1[cur1], nums2[cur2])
			}
			var half = k / 2
			newCur1 := min(cur1+half, len(nums1)) - 1
			newCur2 := min(cur2+half, len(nums2)) - 1
			if nums1[newCur1] <= nums2[newCur2] {
				k -= newCur1 - cur1 + 1
				cur1 = newCur1 + 1
			} else {
				k -= newCur2 - cur2 + 1
				cur2 = newCur2 + 1
			}
		}
	}

	totalLen := len(nums1) + len(nums2)
	if totalLen%2 == 1 {
		return float64(getKthElement((totalLen + 1) / 2))
	} else {
		return float64(getKthElement(totalLen/2)+getKthElement(totalLen/2+1)) / 2
	}
}

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	var getKthElement func(k int) int
	getKthElement = func(k int) int {
		var cur1, cur2 int
		for {
			if cur1 == len(nums1) {
				return nums2[cur2+k-1]
			}
			if cur2 == len(nums2) {
				return nums1[cur1+k-1]
			}
			if k == 1 {
				return min(nums1[cur1], nums2[cur2])
			}
			offset := k / 2
			newCur1 := min(cur1+offset, len(nums1)) - 1
			newCur2 := min(cur2+offset, len(nums2)) - 1
			if nums1[newCur1] <= nums2[newCur2] { // 将cur1 到 newCur1 的元素包含 newCur1 全部排除
				k -= newCur1 + 1 - cur1
				cur1 = newCur1 + 1
			} else {
				k -= newCur2 + 1 - cur2
				cur2 = newCur2 + 1
			}
		}
	}
	totalLen := len(nums1) + len(nums2)
	if totalLen%2 == 0 {
		return float64(getKthElement(totalLen/2)+getKthElement(totalLen/2+1)) / 2
	} else {
		return float64(getKthElement(totalLen/2 + 1))
	}
}

func main() {
	// Example usage
	nums1 := []int{1, 4}
	nums2 := []int{2, 3}
	result := findMedianSortedArrays1(nums1, nums2)
	println(result) // Output: 2.0
}
