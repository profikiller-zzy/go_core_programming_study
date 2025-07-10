package main

import "fmt"

func letterCombinations1(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

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

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var digit2char map[byte]string
	digit2char = map[byte]string{
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
	var dfs func(int)
	dfs = func(index int) {
		if index == len(digits) {
			res = append(res, string(path))
			return
		}
		letter := digit2char[digits[index]]
		for i := 0; i < len(letter); i++ {
			path = append(path, letter[i])
			dfs(index + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}

func main() {
	fmt.Println(letterCombinations("23"))
}
