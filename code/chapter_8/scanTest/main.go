package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

func main() {
	var s scanner.Scanner
	s.Init(strings.NewReader("var x = 42 // example"))

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("Token: %s, Text: %q\n", scanner.TokenString(tok), s.TokenText())
	}
}
