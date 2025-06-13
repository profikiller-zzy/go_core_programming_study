package main

func romanToInt(s string) int {
	romanToIntHash := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var res int
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && romanToIntHash[s[i]] < romanToIntHash[s[i+1]] {
			res -= romanToIntHash[s[i]]
		} else {
			res += romanToIntHash[s[i]]
		}
	}
	return res
}
