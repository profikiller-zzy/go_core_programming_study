package main

import "fmt"

func findAnagrams1(s string, p string) []int {
	res := make([]int, 0)
	if len(s) < len(p) {
		return res
	}

	charHash := make(map[byte]int)
	for index := 0; index < len(p); index++ {
		charHash[p[index]]--
	}
	// 初始化滑动窗口
	for i := 0; i < len(p); i++ {
		charHash[s[i]]++
		if charHash[s[i]] == 0 {
			delete(charHash, s[i])
		}
	}
	if len(charHash) == 0 {
		res = append(res, 0)
	}

	for index := 1; index <= len(s)-len(p); index++ {
		// 滑动窗口向前移动一个字符
		prevChar := s[index-1]
		currChar := s[index+len(p)-1]
		charHash[prevChar]--
		if charHash[prevChar] == 0 {
			delete(charHash, prevChar)
		}
		charHash[currChar]++
		if charHash[currChar] == 0 {
			delete(charHash, currChar)
		}
		if len(charHash) == 0 {
			res = append(res, index)
		}
	}
	return res
}

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return res
	}
	var sChar, pChar [26]int
	for i := 0; i < pLen; i++ {
		sChar[s[i]-'a']++
		pChar[p[i]-'a']++
	}
	if sChar == pChar {
		res = append(res, 0)
	}

	for i := range s[:sLen-pLen] {
		sChar[s[i]-'a']--
		sChar[s[i+pLen]-'a']++
		if sChar == pChar {
			res = append(res, i+1)
		}
	}
	return res
}

func main() {
	fmt.Println(findAnagrams("baa", "aa"))
}
