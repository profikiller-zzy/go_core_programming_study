package IsPalindrome

import "sort"

// 练习 7.10： sort.Interface类型也可以适用在其它地方。编写一个IsPalindrome(s sort.Interface) bool函数表明序列s是否是回文序列，
// 换句话说反向排序不会改变这个序列。假设如果!s.Less(i, j) && !s.Less(j, i)则索引i和j上的元素相等。

func IsPalindrome(s sort.Interface) bool {
	var len = s.Len()
	for i := 0; i < len/2+1; i++ {
		j := len - i
		if s.Less(i, j) || s.Less(j, i) { // 这两个元素不相等就不知回文串
			return false
		}
	}
	return true
}
