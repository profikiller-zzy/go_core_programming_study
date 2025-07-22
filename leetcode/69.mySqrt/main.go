package main

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	left, right := 1, x/2
	for left <= right {
		mid := (left + right) / 2
		power := mid * mid
		if power == x {
			return mid
		} else if power < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left - 1
}

func main() {
	// Example usage
	result := mySqrt(8)
	println(result) // Output: 2
}
