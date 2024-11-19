package eval

import (
	"fmt"
	"strings"
)

// 练习 7.14： 定义一个新的满足Expr接口的具体类型并且提供一个新的操作例如对它运算单元中的最小值的计算。
// 因为Parse函数不会创建这个新类型的实例，为了使用它你可能需要直接构造一个语法树（或者继承parser接口）。

// 定义 min 类型，包含多个子表达式
type min struct {
	operands []Expr
}

// 实现 Expr 接口的 String 方法
func (m min) String() string {
	// 生成 min 函数形式的字符串表示
	var parts []string
	for _, operand := range m.operands {
		parts = append(parts, operand.String())
	}
	return fmt.Sprintf("min(%s)", strings.Join(parts, ", "))
}

// Eval 实现 Eval 方法，用于计算表达式列表中的最小值
func (m min) Eval(env Env) float64 {
	if len(m.operands) == 0 {
		return 0 // or handle as an error, depending on design choice
	}
	minVal := m.operands[0].Eval(env)
	for _, operand := range m.operands[1:] {
		val := operand.Eval(env)
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}

func (m min) Check(vars map[Var]bool) error {
	for _, operand := range m.operands {
		err := operand.Check(vars)
		if err != nil { // 返回检查出来的第一个不合法的表达式
			return err
		}
	}
	return nil
}
