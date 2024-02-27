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
    // Comment begins at column 5.

This line should not be included in the output.

/*
This multiline comment
should be extracted in
its entirety.
*/
`

	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	s.Filename = "comments"
	s.Mode ^= scanner.SkipComments // don't skip comments

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		txt := s.TokenText()
		if strings.HasPrefix(txt, "//") || strings.HasPrefix(txt, "/*") {
			fmt.Printf("%s: %s\n", s.Position, txt)
		}
	}
}
