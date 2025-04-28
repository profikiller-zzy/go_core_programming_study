package main

type Stack struct {
	len  int
	data []byte
}

func newStack() *Stack {
	return &Stack{
		len:  0,
		data: make([]byte, 0),
	}
}

func (s *Stack) push(c byte) {
	s.data = append(s.data, c)
	s.len++
}

func (s *Stack) pop() byte {
	if s.len == 0 {
		return ' '
	}
	c := s.data[s.len-1]
	s.data = s.data[:s.len-1]
	s.len--
	return c
}

func isValid(s string) bool {
	stack := newStack()
	for index := 0; index < len(s); index++ {
		switch s[index] {
		case ')':
			ch := stack.pop()
			if ch != '(' {
				return false
			}
		case ']':
			ch := stack.pop()
			if ch != '[' {
				return false
			}
		case '}':
			ch := stack.pop()
			if ch != '{' {
				return false
			}
		default:
			stack.push(s[index])
		}
	}
	if stack.len != 0 {
		return false
	}
	return true
}
