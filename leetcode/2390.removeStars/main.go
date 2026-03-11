package main

func removeStars(s string) string {
	byteStack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch == '*' {
			byteStack = byteStack[:len(byteStack)-1]
		} else {
			byteStack = append(byteStack, ch)
		}
	}
	return string(byteStack)
}
