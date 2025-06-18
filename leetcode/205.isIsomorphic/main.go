package main

func isIsomorphic(s string, t string) bool {
	sCharHash, tCharHash := map[byte]byte{}, map[byte]byte{}
	for index := 0; index < len(s); index++ {
		sChar, tChar := s[index], t[index]
		if sCharHash[sChar] > 0 && sCharHash[sChar] != tChar || tCharHash[tChar] > 0 && tCharHash[tChar] != sChar {
			return false
		}
		sCharHash[sChar] = tChar
		tCharHash[tChar] = sChar
	}
	return true
}
