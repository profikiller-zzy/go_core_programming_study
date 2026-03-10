package main

import "fmt"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	digit2char := map[byte]string{
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
	var dfs func(index int)
	dfs = func(index int) {
		if len(path) == len(digits) {
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
