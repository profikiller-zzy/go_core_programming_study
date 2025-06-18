package main

func canConstruct(ransomNote string, magazine string) bool {
	providedChar := map[byte]int{}
	for index := 0; index < len(magazine); index++ {
		providedChar[magazine[index]]++
	}
	for index := 0; index < len(ransomNote); index++ {
		providedChar[ransomNote[index]]--
	}
	for _, v := range providedChar {
		if v < 0 {
			return false
		}
	}
	return true
}
