package main

import "fmt"

// https://leetcode.cn/problems/group-anagrams/description/

func groupAnagrams(strs []string) [][]string {
	var res [][]string = make([][]string, 0)
	if len(strs) == 0 {
		return res
	} else if len(strs) == 1 {
		res = append(res, []string{strs[0]})
		return res
	}

	// processStr 将这个字符串重新排序，如果是res，排序之后就是ers
	processStr := func(src string) string {
		srcBytes := []byte(src)
		for i := 0; i < len(srcBytes)-1; i++ {
			swapped := false
			for j := 0; j < len(srcBytes)-1-i; j++ {
				if srcBytes[j] > srcBytes[j+1] {
					srcBytes[j], srcBytes[j+1] = srcBytes[j+1], srcBytes[j]
					swapped = true
				}
			}
			if !swapped {
				break
			}
		}
		return string(srcBytes)
	}

	// 字母异位词重新排序后的字符串都是相同的，直接将排序后的字符串作为key存入map中
	strHash := make(map[string][]string)
	for _, str := range strs {
		strHash[processStr(str)] = append(strHash[processStr(str)], str)
	}
	for _, v := range strHash {
		res = append(res, v)
	}
	return res
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
