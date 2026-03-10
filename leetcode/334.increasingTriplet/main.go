package main

func increasingTriplet(nums []int) bool {
	tails := make([]int, 0, len(nums))
	for _, num := range nums {
		if len(tails) == 0 || num > tails[len(tails)-1] {
			tails = append(tails, num)
		} else {
			left, right := 0, len(tails)-1
			for left <= right {
				mid := (left + right) / 2
				if num > tails[mid] {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
			tails[left] = num
		}
	}
	return len(tails) >= 3
}
