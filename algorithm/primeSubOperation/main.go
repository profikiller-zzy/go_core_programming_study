package main

import "fmt"

var p = []int{0} // 哨兵，避免二分越界

// init 通过筛质数来获得质数列表
func init() {
	const mx = 1000
	np := [mx]bool{}
	for i := 2; i < mx; i++ {
		if !np[i] {
			p = append(p, i) // 预处理质数列表
			for j := i * i; j < mx; j += i {
				np[j] = true
			}
		}
	}
}

func main() {
	nums := []int{998, 2}
	fmt.Println(primeSubOperation(nums))
}

func primeSubOperation(nums []int) bool {
	// 质数数组，存储小于数组nums中最大数的质数数组
	primeArray := make([]int, 0)

	// isPrime 判断一个数是不是质数
	isPrime := func(n int) bool {
		if n <= 3 {
			return n > 1
		}
		//当n>3时，质数无法被比它小的数整除
		for i := 2; i < n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	// 求最大值
	maxVal := nums[0]
	for _, v := range nums[1:] {
		if v > maxVal {
			maxVal = v
		}
	}

	for prime := 2; prime < maxVal; prime++ {
		if isPrime(prime) {
			primeArray = append(primeArray, prime)
		}
	}

	// 假如没有一个质数严格小于数组里面的每一个数
	if len(primeArray) == 0 {
		for i := 1; i < len(nums); i++ {
			if nums[i] <= nums[i-1] {
				return false
			}
		}
		return true
	}

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			// 对于数组的第一个数，直接让这个数减去小于这个数的最大质数
			index := 0
			for index < len(primeArray) && primeArray[index] < nums[i] {
				index++
			}
			if index != 0 {
				nums[i] = nums[i] - primeArray[index-1]
			}
		} else {
			// 对于数组的第二个及之后的数，让这个数减去小于这个数并且结果还大于前一个数的最大质数
			// 如果这个质数不存在，直接返回false
			index := 0
			for index < len(primeArray) && primeArray[index] < nums[i] && nums[i]-primeArray[index] > nums[i-1] {
				index++
			}
			if nums[i] > nums[i-1] && index != 0 {
				nums[i] = nums[i] - primeArray[index-1]
			}
		}

		if i > 0 && nums[i] <= nums[i-1] {
			return false
		}
	}
	//fmt.Println(nums)
	return true
}
