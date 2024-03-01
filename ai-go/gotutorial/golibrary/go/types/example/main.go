/*
@File   : main.go
@Author : pan
@Time   : 2024-03-01 10:48:47
*/
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	// Parse the source files for a package.
	fset := token.NewFileSet()
	var files []*ast.File
	for _, src := range []string{
		`package main
import "fmt"
func main() {
    fmt.Println("Hello, world!")
}`,
		`package main
func foo() int {
    return 42
}`,
	} {
		f := mustParse(fset, src)
		files = append(files, f)
	}

	// Print the tree of scopes for the package.
	for _, f := range files {
		printScope(fset, f, 0)
		fmt.Println()
	}
}

func mustParse(fset *token.FileSet, input string) *ast.File {
	f, err := parser.ParseFile(fset, "<input>", input, 0)
	if err != nil {
		panic(err)
	}
	return f
}

func printScope(fset *token.FileSet, n ast.Node, depth int) {
	switch n := n.(type) {
	case *ast.File:
		printFileScope(fset, n, depth)
	case *ast.BlockStmt:
		printBlockScope(fset, n, depth)
	case *ast.FuncDecl:
		printFuncScope(fset, n, depth)
	default:
		fmt.Printf("%s<%T>\n", indent(depth), n)
	}
}

func printFileScope(fset *token.FileSet, f *ast.File, depth int) {
	fmt.Printf("%sPackage %s\n", indent(depth), f.Name.Name)
	for _, decl := range f.Decls {
		printScope(fset, decl, depth+1)
	}
}

func printBlockScope(fset *token.FileSet, b *ast.BlockStmt, depth int) {
	fmt.Printf("%sBlock\n", indent(depth))
	for _, stmt := range b.List {
		printScope(fset, stmt, depth+1)
	}
}

func printFuncScope(fset *token.FileSet, f *ast.FuncDecl, depth int) {
	fmt.Printf("%sFunction %s\n", indent(depth), f.Name.Name)
	printScope(fset, f.Type, depth+1)
	printScope(fset, f.Body, depth+1)
}

func indent(depth int) string {
	return string(make([]byte, depth*2))
}
