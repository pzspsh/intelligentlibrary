/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 17:48:53
*/
package main

import (
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
)

func main() {
	// src and test are two source files that make up
	// a package whose documentation will be computed.
	const src = `
// This is the package comment.
package p

import "fmt"

// This comment is associated with the Greet function.
func Greet(who string) {
	fmt.Printf("Hello, %s!\n", who)
}
`
	const test = `
package p_test

// This comment is associated with the ExampleGreet_world example.
func ExampleGreet_world() {
	Greet("world")
}
`

	// Create the AST by parsing src and test.
	fset := token.NewFileSet()
	files := []*ast.File{
		mustParse(fset, "src.go", src),
		mustParse(fset, "src_test.go", test),
	}

	// Compute package documentation with examples.
	p, err := doc.NewFromFiles(fset, files, "example.com/p")
	if err != nil {
		panic(err)
	}

	fmt.Printf("package %s - %s", p.Name, p.Doc)
	fmt.Printf("func %s - %s", p.Funcs[0].Name, p.Funcs[0].Doc)
	fmt.Printf(" ⤷ example with suffix %q - %s", p.Funcs[0].Examples[0].Suffix, p.Funcs[0].Examples[0].Doc)

}

func mustParse(fset *token.FileSet, filename, src string) *ast.File {
	f, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	return f
}
