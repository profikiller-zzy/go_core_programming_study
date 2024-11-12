package main

import (
	"GoStudy/Answer/ch7/exprEval/eval"
	"fmt"
)

func main() {
	exprStr := "sqrt(A / pi)"
	parsedExpr, err := eval.Parse(exprStr)
	if err == nil {
		err = parsedExpr.Check(map[eval.Var]bool{})
	}
	if err != nil {
		fmt.Println("Error parsing expression:", err)
		return
	}

	exprStr1 := parsedExpr.String()
	fmt.Println(exprStr1)

	// 重新解析 exprStr1，并检查它是否生成相同的语法树
	parsedExpr1, err := eval.Parse(exprStr1)
	if err == nil {
		err = parsedExpr1.Check(map[eval.Var]bool{})
	}
	if err != nil {
		fmt.Println("Error parsing re-parsed expression:", err)
		return
	}

	exprStr2 := parsedExpr1.String()
	fmt.Println(exprStr2)

	// 比较两次生成的字符串表示是否相同
	if exprStr1 == exprStr2 {
		fmt.Println("The two parsed expressions are identical.")
	} else {
		fmt.Println("The two parsed expressions are different.")
	}
}
