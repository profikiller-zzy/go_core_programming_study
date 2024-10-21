package main

import "fmt"

func letterCombinations(digits string) []string {
	var digitLetter map[byte]string
	digitLetter = map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var res []string
	var path []byte
	if len(digits) == 0 {
		return res
	}

	// 回溯函数，用来构建所有可能的字母组合
	var backtrack func(index int)
	backtrack = func(index int) {
		if index == len(digits) {
			res = append(res, string(path))
			return
		}

		letters := digitLetter[digits[index]] // 当前数字代表的号码

		for i := 0; i < len(letters); i++ {
			path = append(path, letters[i]) // 先选定当前号码中的一个字母
			backtrack(index + 1)            // 选定下一个号码
			path = path[:len(path)-1]       // 回溯，移除刚刚选定的字母
		}
	}
	backtrack(0) // 从开头开始回溯
	return res
}

func main() {
	fmt.Println(letterCombinations("23"))
}
