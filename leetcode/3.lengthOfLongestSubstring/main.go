package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var curCharSet = make(map[byte]int)
	maxArea := 0
	left, right := 0, 0
	for index := 0; index < len(s); index++ {
		if _, ok := curCharSet[s[index]]; !ok { // 当前滑动窗口中没有这个字符
			right = index
			curCharSet[s[index]] = index
			curArea := right - left + 1
			if curArea > maxArea {
				maxArea = curArea
			}
		} else { // 当前滑动窗口中已经有这个字符了
			// 将滑动窗口起始到这个字符的位置移出
			thisCharIndex := curCharSet[s[index]]
			for i := left; i <= thisCharIndex-1; i++ {
				delete(curCharSet, s[i])
			}
			curCharSet[s[index]] = index
			left = thisCharIndex + 1
		}
	}
	return maxArea
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
