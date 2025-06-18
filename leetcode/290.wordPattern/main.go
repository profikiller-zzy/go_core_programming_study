package wordPattern

import "strings"

func wordPattern(pattern string, s string) bool {
	charToWord, wordToChar := map[byte]string{}, map[string]byte{}
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	for index := 0; index < len(pattern); index++ {
		ch, word := pattern[index], words[index]
		if charToWord[ch] != "" && charToWord[ch] != word || wordToChar[word] != 0 && wordToChar[word] != ch {
			return false
		}
		charToWord[ch] = word
		wordToChar[word] = ch
	}
	return true
}
