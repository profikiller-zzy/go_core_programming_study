package main

import "fmt"

func lengthOfLongestSubstring1(s string) int {
	var curCharSet = make(map[byte]int)
	maxArea := 0
	left, right := 0, 0
	for index := 0; index < len(s); index++ {
		ch := s[index]
		if _, ok := curCharSet[s[index]]; !ok { // 当前滑动窗口中没有这个字符
			right = index
			curCharSet[ch] = index
			curArea := right - left + 1
			if curArea > maxArea {
				maxArea = curArea
			}
		} else {
			// 当前滑动窗口中已经有这个字符了
			// 将滑动窗口起始到这个字符的位置移出(可以直接移出就是因为他们的数量一定是1)
			thisCharIndex := curCharSet[ch]
			for i := left; i <= thisCharIndex-1; i++ {
				delete(curCharSet, s[i])
			}
			curCharSet[ch] = index
			left = thisCharIndex + 1
		}
	}
	return maxArea
}

func lengthOfLongestSubstring(s string) int {
	window := make(map[rune]struct{})
	var left, result int
	for right := 0; right < len(s); right++ {
		ch := rune(s[right])
		if _, ok := window[ch]; !ok {
			window[ch] = struct{}{}
			if result < len(window) {
				result = len(window)
			}
		} else {
			for ; left < right && ch != rune(s[left]); left++ {
				delete(window, rune(s[left]))
			}
			left++
		}
	}
	return result
}

func main() {
	s := "aabaab!bb"
	fmt.Println(lengthOfLongestSubstring(s))
}
