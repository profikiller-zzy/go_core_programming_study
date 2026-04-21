package main

import "fmt"

func removeDuplicateLetters(s string) string {
	var stack []byte
	var byteNum [26]int
	var inResult [26]bool
	for i := 0; i < len(s); i++ {
		byteNum[int(s[i]-'a')]++
	}
	for i := 0; i < len(s); i++ {
		ch := s[i]
		byteNum[int(ch-'a')]--
		if inResult[int(ch-'a')] {
			continue
		}
		for len(stack) > 0 && ch < stack[len(stack)-1] && byteNum[int(stack[len(stack)-1]-'a')] > 0 {
			// 当前字符的字典序小于栈顶字符的字典序、并且栈顶字符在右边还会出现
			inResult[int(stack[len(stack)-1]-'a')] = false
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, ch)
		inResult[int(ch-'a')] = true
	}
	return string(stack)
}

func main() {
	fmt.Println(removeDuplicateLetters("bbcaac"))
}
