package _5_9

import "strings"

// 练习 5.9： 编写函数Expand，将s中的"foo"替换为f("foo")的返回值。

func Expand(s string, f func(string) string) string {
	return strings.Replace(s, "foo", f("foo"), -1)
}

func F(str string) string {
	if str == "foo" {
		return "oof"
	}
	return ""
}
