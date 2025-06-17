package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	need := make(map[byte]int)   // 存储t中每个字符的需求数量
	window := make(map[byte]int) // 存储当前滑动窗口中的字符数量

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left, right := 0, 0
	valid := 0 // 满足条件的字符数量

	start, minLen := 0, len(s)+1

	for ; right <= len(s)-1; right++ {
		c := s[right]

		if _, ok := need[c]; ok { // 当前字符在t中
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		// 判断左侧窗口是否要收缩
		for valid == len(need) {
			// 更新最小覆盖子串
			if right-left < minLen {
				start = left
				minLen = right - left
			}

			d := s[left]
			left++

			if _, ok := need[d]; ok { // left所指的字符在t中
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if minLen == len(s)+1 {
		return ""
	}
	return s[start : start+minLen+1]
}

func minWindow1(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	tCharHash, windowCharHash := make(map[byte]int), make(map[byte]int)
	for index := 0; index < len(t); index++ {
		tCharHash[t[index]]++
	}
	var isValid func() bool
	isValid = func() bool {
		for k, v := range tCharHash {
			if windowCharHash[k] < v {
				return false
			}
		}
		return true
	}

	sLen := len(s)
	minLen := math.MaxInt32
	resLeft, resRight := -1, -1
	for left, right := 0, 0; right < sLen; right++ {
		if tCharHash[s[right]] > 0 { // 没有越界并且当前的字符在t中
			windowCharHash[s[right]]++
		}
		for isValid() && left <= right { // 收缩左边界
			if right-left+1 < minLen {
				resLeft = left
				resRight = right
				minLen = right - left + 1
			}
			if tCharHash[s[left]] > 0 {
				windowCharHash[s[left]]--
			}
			left++
		}
	}
	if resLeft == -1 {
		return ""
	}
	return s[resLeft : resRight+1]
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t)) // Output: "BANC"
}
