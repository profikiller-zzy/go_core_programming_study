package main

import (
	"errors"
	"fmt"
	"math"
)

// 练习5.15： 编写类似sum的可变参数函数max和min。考虑不传参时，max和min该如何处理，再编写至少接收1个参数的版本。

// max 接收类型为int的可变参数，返回它们的最大值。当函数没有传参时返回错误
func max(nums ...int) (int, error) {
	err := errors.New("参数个数为零")
	if len(nums) == 0 {
		return 0, err
	}
	// 初始化最大值为最小数
	maxNum := math.MinInt
	for _, val := range nums {
		if val > maxNum {
			maxNum = val
		}
	}
	return maxNum, nil
}

func main() {
	var maxNum int
	// 可以选择处理错误, 也可以选择不处理，直接丢弃掉
	maxNum, err := max(4, 2, 3, 9)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(maxNum)
}
