package main

import (
	"go_core_programming/Answer/ch7/7_9_Expr/webCal"
)

func main() {
	// 创建一个 webCalculator 实例
	wc := &webCal.WebCalculator{}
	wc.ListenAndServe()
}
