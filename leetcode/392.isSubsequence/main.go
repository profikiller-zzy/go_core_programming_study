package main

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	var sIndex, tIndex int
	for tIndex < len(t) {
		if s[sIndex] == t[tIndex] {
			sIndex++
			tIndex++
			if sIndex == len(s) {
				return true
			}
			continue
		}
		tIndex++
	}
	return false
}

func main() {
	s := "abc"
	t := "ahbgdc"
	result := isSubsequence(s, t)
	println(result) // Output: true
}
