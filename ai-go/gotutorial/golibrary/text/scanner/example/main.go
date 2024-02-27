/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 23:49:40
*/
package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

func main() {
	const src = `
// This is scanned code.
if a > 10 {
	someParsable = text
}`

	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	s.Filename = "example"
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
	}

}
