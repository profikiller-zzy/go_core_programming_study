package main

func isValid(s string) bool {
	var stack []byte
	for index := 0; index < len(s); index++ {
		switch s[index] {
		case ')':
			if len(stack) <= 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) <= 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) <= 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			stack = append(stack, s[index])
		}
	}
	return len(stack) == 0
}
