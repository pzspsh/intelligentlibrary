/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 17:23:59
*/
package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

/*
这个例子显示了AST在打印用于调试时的样子。
*/
func main() {
	// src is the input for which we want to print the AST.
	src := `
package main
func main() {
	println("Hello, World!")
}
`

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	// Print the AST.
	ast.Print(fset, f)

}
