// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package eval

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

// ---- lexer ----

// This lexer is similar to the one described in Chapter 13.
// lexer 是一个用于词法分析的结构体，包含了`scanner.Scanner`实例和一个当前的标记
type lexer struct {
	scan  scanner.Scanner
	token rune // current lookahead token
}

// next 获取下一个token
func (lex *lexer) next() { lex.token = lex.scan.Scan() }

// text 返回当前token的文本
func (lex *lexer) text() string { return lex.scan.TokenText() }

// lexPanic 是用于词法分析错误的自定义类型，当解析过程中遇到无法解析的token时，会使用`panic`抛出`lexPanic`异常
type lexPanic string

// describe returns a string describing the current token, for use in errors.
func (lex *lexer) describe() string {
	switch lex.token {
	case scanner.EOF:
		return "end of file"
	case scanner.Ident: // 标识符
		return fmt.Sprintf("identifier %s", lex.text())
	case scanner.Int, scanner.Float:
		return fmt.Sprintf("number %s", lex.text())
	}
	return fmt.Sprintf("%q", rune(lex.token)) // any other rune
}

func precedence(op rune) int {
	switch op {
	case '*', '/':
		return 2
	case '+', '-':
		return 1
	}
	return 0
}

// ---- parser ----

// Parse parses the input string as an arithmetic expression.
//
//	expr = num                         a literal number, e.g., 3.14159
//	     | id                          a variable name, e.g., x
//	     | id '(' expr ',' ... ')'     a function call
//	     | '-' expr                    a unary operator (+-)
//	     | expr '+' expr               a binary operator (+-*/)
func Parse(input string) (_ Expr, err error) {
	defer func() {
		switch x := recover().(type) {
		case nil:
			// no panic
		case lexPanic:
			err = fmt.Errorf("%s", x)
		default:
			// unexpected panic: resume state of panic.
			panic(x)
		}
	}()
	// 初始化扫描器
	lex := new(lexer)
	lex.scan.Init(strings.NewReader(input))
	// `ScanIdents`: 启用标识符扫描模式，用于识别标识符（例如变量名、函数名等）
	// `ScanInts`: 表示启用整数扫描模式，用于识别整数常量
	// `ScanFloats`: 表示启用浮点数扫描模式，用于识别浮点数常量
	lex.scan.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats
	lex.next() // initial lookahead
	e := parseExpr(lex)
	if lex.token != scanner.EOF {
		return nil, fmt.Errorf("unexpected %s", lex.describe())
	}
	return e, nil
}

// parseExpr 解析表达式
func parseExpr(lex *lexer) Expr { return parseBinary(lex, 1) }

// binary = unary ('+' binary)*
// parseBinary stops when it encounters an
// operator of lower precedence than prec1.
func parseBinary(lex *lexer, prec1 int) Expr {
	// 调用parseUnary解析左边的表达式
	lhs := parseUnary(lex)
	for prec := precedence(lex.token); prec >= prec1; prec-- {
		for precedence(lex.token) == prec { // 确保了在同一优先级的所有运算符被处理
			op := lex.token // 将当前的运算符存储在变量 op
			lex.next()      // consume operator
			rhs := parseBinary(lex, prec+1)
			lhs = binary{op, lhs, rhs}
		}
	}
	return lhs
}

// unary = '+' expr | primary
// parseUnary 解析一个一元表达式
func parseUnary(lex *lexer) Expr {
	if lex.token == '+' || lex.token == '-' {
		op := lex.token
		lex.next() // consume '+' or '-'
		return unary{op, parseUnary(lex)}
	}
	return parsePrimary(lex)
}

// primary = id
//
//	| id '(' expr ',' ... ',' expr ')'
//	| num
//	| '(' expr ')'
func parsePrimary(lex *lexer) Expr {
	switch lex.token {
	case scanner.Ident: // 当前token是一个标识符（`scanner.Ident`），则它可能是一个变量名或者是函数调用
		id := lex.text()
		lex.next()            // consume Ident
		if lex.token != '(' { // 下一个token不是`(`，则它是一个变量名
			return Var(id)
		}
		lex.next()      // consume '(' ，将其视为函数调用，解析这个函数调用
		var args []Expr // 函数调用的参数参数列表
		if lex.token != ')' {
			for {
				args = append(args, parseExpr(lex))
				if lex.token != ',' {
					break
				}
				lex.next() // consume ','
			}
			if lex.token != ')' {
				msg := fmt.Sprintf("got %s, want ')'", lex.describe())
				panic(lexPanic(msg))
			}
		}
		lex.next() // consume ')'
		return call{id, args}

	case scanner.Int, scanner.Float:
		f, err := strconv.ParseFloat(lex.text(), 64)
		if err != nil {
			panic(lexPanic(err.Error()))
		}
		lex.next() // consume number
		return literal(f)

	case '(':
		lex.next() // consume '('
		e := parseExpr(lex)
		if lex.token != ')' {
			msg := fmt.Sprintf("got %s, want ')'", lex.describe())
			panic(lexPanic(msg))
		}
		lex.next() // consume ')'
		return e
	}
	msg := fmt.Sprintf("unexpected %s", lex.describe())
	panic(lexPanic(msg))
}
