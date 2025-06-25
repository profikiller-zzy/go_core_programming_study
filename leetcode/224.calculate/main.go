package main

import (
	"container/list"
	"fmt"
	"strconv"
	"unicode"
)

func calculate1(s string) int {
	result, stack, sign := 0, list.New(), 1
	for index := 0; index < len(s); {
		if s[index] == ' ' {
			index++
		} else if s[index] == '+' {
			sign = 1
			index++
		} else if s[index] == '-' {
			sign = -1
			index++
		} else if s[index] <= '9' && s[index] >= '0' {
			val := int(s[index] - '0')
			for index+1 < len(s) && s[index+1] <= '9' && s[index+1] >= '0' {
				index++
				val = val*10 + int(s[index]-'0')
			}
			result += val * sign
			index++
		} else if s[index] == '(' {
			stack.PushBack(result)
			stack.PushBack(sign)
			index++
			// 重置resul和sign，现在开始计算括号内的值
			result = 0
			sign = 1
		} else if s[index] == ')' {
			// 括号结束之后，栈顶和次栈顶分别是上一个sign和result
			result = stack.Remove(stack.Back()).(int)*result + stack.Remove(stack.Back()).(int)
			index++
		}
	}
	return result
}

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: make([]T, 0)}
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val, true
}

func (s *Stack[T]) Top() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}
	return s.data[len(s.data)-1], true
}

func (s *Stack[T]) Empty() bool {
	return len(s.data) == 0
}

// tokenize 表达式分词 + 支持一元负号
func tokenize(expr string) []string {
	var tokens []string
	var index int
	n := len(expr)
	for index < n {
		ch := expr[index]

		if ch == ' ' {
			index++
			continue
		}

		// 负号处理
		if ch == '-' {
			if index == 0 || (index > 0 && (expr[index-1] == '(' || expr[index-1] == '+' || expr[index-1] == '-' || expr[index-1] == '*' || expr[index-1] == '/')) {
				// 是一元负号，尝试读取负数
				j := index + 1
				for j < n && unicode.IsDigit(rune(expr[j])) {
					j++
				}
				tokens = append(tokens, expr[index:j])
				index = j
				continue
			}
		}

		if unicode.IsDigit(rune(ch)) {
			j := index
			for j < n && unicode.IsDigit(rune(expr[j])) {
				j++
			}
			tokens = append(tokens, expr[index:j])
			index = j
			continue
		}

		// 运算符或括号
		tokens = append(tokens, string(ch))
		index++
	}
	return tokens
}

// infixToPostfix 中缀转后缀表达式
func infixToPostfix(tokens []string) []string {
	var output []string
	opStack := NewStack[string]()

	precedence := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			for {
				top, ok := opStack.Top()
				if !ok || top == "(" {
					break
				}
				if precedence[top] >= precedence[token] {
					op, _ := opStack.Pop()
					output = append(output, op)
				} else {
					break
				}
			}
			opStack.Push(token)

		case "(":
			opStack.Push(token)

		case ")":
			for {
				op, _ := opStack.Pop()
				if op == "(" {
					break
				}
				output = append(output, op)
			}

		default: // 数字
			output = append(output, token)
		}
	}

	// 剩余操作符出栈
	for !opStack.Empty() {
		op, _ := opStack.Pop()
		output = append(output, op)
	}

	return output
}

// evalPostfix 计算后缀表达式
func evalPostfix(postfix []string) int {
	numStack := NewStack[int]()

	for _, token := range postfix {
		switch token {
		case "+", "-", "*", "/":
			b, _ := numStack.Pop()
			a, _ := numStack.Pop()
			switch token {
			case "+":
				numStack.Push(a + b)
			case "-":
				numStack.Push(a - b)
			case "*":
				numStack.Push(a * b)
			case "/":
				numStack.Push(a / b) // 默认整数除法
			}
		default:
			num, _ := strconv.Atoi(token)
			numStack.Push(num)
		}
	}
	result, _ := numStack.Pop()
	return result
}

func calculateExpression(expr string) int {
	tokens := tokenize(expr)
	postfix := infixToPostfix(tokens)
	return evalPostfix(postfix)
}

func main() {
	s := "- (3 + (4 + 5))"
	fmt.Println(calculate1(s))
}
