package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	returnLowestLen := func() int {
		minLen := len(strs[0])
		for _, str := range strs {
			if len(str) < minLen {
				minLen = len(str)
			}
		}
		return minLen
	}
	minLen := returnLowestLen() // 获取最短字符串的长度
	resStr := ""                // 用于存储最长公共前缀

	// 遍历每个字符的索引，最多到最短字符串的长度
	for i := 0; i < minLen; i++ {
		ch := strs[0][i] // 取第一个字符串的第i个字符
		for _, str := range strs {
			if str[i] != ch { // 如果有一个字符串的第i个字符不匹配
				return resStr // 直接返回结果
			}
		}
		resStr += string(ch) // 将公共字符加到结果中
	}

	return resStr // 返回最终的最长公共前缀
}

func main() {

}
