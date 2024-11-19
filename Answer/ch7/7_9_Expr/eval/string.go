package eval

import (
	"fmt"
	"strconv"
	"strings"
)

// 练习 7.13： 为Expr增加一个String方法来打印美观的语法树。当再一次解析的时候，检查它的结果是否生成相同的语法树。

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return strconv.FormatFloat(float64(l), 'f', -1, 64)
}

func (u unary) String() string {
	return fmt.Sprintf("(%c%s)", u.op, u.x.String())
}

func (b binary) String() string {
	return fmt.Sprintf("(%s %c %s)", b.x.String(), b.op, b.y.String())
}

func (c call) String() string {
	var args = make([]string, len(c.args))
	for i, arg := range c.args {
		args[i] = arg.String()
	}
	return fmt.Sprintf("%s(%s)", c.fn, strings.Join(args, ", "))
}
