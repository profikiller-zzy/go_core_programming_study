package main

import "fmt"

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	if len(s) < len(p) {
		return res
	}

	charHash := make(map[rune]int)
	for _, char := range p {
		charHash[char]--
	}
	// 初始化滑动窗口
	for i := 0; i < len(p); i++ {
		charHash[rune(s[i])]++
		if charHash[rune(s[i])] == 0 {
			delete(charHash, rune(s[i]))
		}
	}
	if len(charHash) == 0 {
		res = append(res, 0)
	}

	for index := 1; index <= len(s)-len(p); index++ {
		// 滑动窗口向前移动一个字符
		prevChar := rune(s[index-1])
		currChar := rune(s[index+len(p)-1])
		charHash[prevChar]--
		if charHash[prevChar] == 0 {
			delete(charHash, rune(prevChar))
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

func main() {
	fmt.Println(findAnagrams("baa", "aa"))
}
