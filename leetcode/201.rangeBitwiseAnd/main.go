package main

// rangeBitwiseAnd 计算给定范围内所有整数的按位与结果。
// 如果你观察多个连续数的按位与，会发现：
//
// ✅ 只保留 left 和 right 的公共前缀部分，其余全部变成了 0。
//
// 所以我们：
//  1. 不断右移 left 和 right，直到它们相等；
//  2. 记录右移的次数 shift；
//  3. 最后将相等的值左移回来。
func rangeBitwiseAnd(left int, right int) int {
	shift := 0
	for left < right {
		left >>= 1
		right >>= 1
		shift++
	}
	return left << shift
}
