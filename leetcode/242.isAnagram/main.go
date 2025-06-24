package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sCharHash, tCharHash := map[byte]int{}, map[byte]int{}
	for index := 0; index < len(s); index++ {
		sCharHash[s[index]]++
		tCharHash[t[index]]++
	}

	for k, v := range sCharHash {
		if tCharHash[k] != v {
			return false
		}
	}
	return true
}
