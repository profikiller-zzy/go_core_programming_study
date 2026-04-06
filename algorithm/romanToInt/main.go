package main

import "fmt"

func romanToInt(s string) int {
	romanMap := make(map[string]int)
	romanMap = map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}
	res := 0
	for i := 0; i < len(s); {
		if i < len(s)-1 {
			if value, exist := romanMap[string(s[i:i+2])]; exist {
				res += value
				i += 2
			} else {
				value, _ := romanMap[string(s[i:i+1])]
				res += value
				i += 1
			}
		} else {
			value, _ := romanMap[string(s[i:i+1])]
			res += value
			i += 1
		}
	}
	return res
}

func main() {
	fmt.Println(romanToInt("III"))
}
