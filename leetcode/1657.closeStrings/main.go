package main

import (
	"reflect"
	"sort"
)

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	charCount1, charCount2 := make([]int, 26), make([]int, 26)
	for i := 0; i < len(word1); i++ {
		charCount1[word1[i]-'a']++
		charCount2[word2[i]-'a']++
	}
	// 比较字符集是否相同
	for i := 0; i < 26; i++ {
		if charCount1[i] > 0 && charCount2[i] == 0 || charCount1[i] == 0 && charCount2[i] > 0 {
			return false
		}
	}
	sort.Ints(charCount1)
	sort.Ints(charCount2)
	// 比较字符的数量集是否相同
	return reflect.DeepEqual(charCount1, charCount2)
}
