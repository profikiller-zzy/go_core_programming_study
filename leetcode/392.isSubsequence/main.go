package main

func isSubsequence(s string, t string) bool {
	var i, j int
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}

func main() {
	s := "abc"
	t := "ahbgdc"
	result := isSubsequence(s, t)
	println(result) // Output: true
}
