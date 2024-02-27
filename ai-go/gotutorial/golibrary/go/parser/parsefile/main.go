/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 17:53:31
*/
package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet() // positions are relative to fset

	src := `package foo

import (
	"fmt"
	"time"
)

func bar() {
	fmt.Println(time.Now())
}`

	// Parse src but stop after processing the imports.
	f, err := parser.ParseFile(fset, "", src, parser.ImportsOnly)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the imports from the file's AST.
	for _, s := range f.Imports {
		fmt.Println(s.Path.Value)
	}

}
