package main

func largestAltitude(gain []int) int {
	maxHeight, curHeight := 0, 0
	for _, g := range gain {
		curHeight += g
		if curHeight > maxHeight {
			maxHeight = curHeight
		}
	}
	return maxHeight
}
