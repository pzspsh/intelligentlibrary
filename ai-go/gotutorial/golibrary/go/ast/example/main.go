/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 17:28:35
*/
package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
)

func main() {
	fset := token.NewFileSet()
	src := `package main

import (
	"context"
)

// Foo 结构体
type Foo struct {
	i int
}

// Bar 接口
type Bar interface {
	Do(ctx context.Context) error
}

// main方法
func main() {
    a := 1
}`
	// 这里取绝对路径，方便打印出来的语法树可以转跳到编辑器
	path, _ := filepath.Abs(src)
	f, err := parser.ParseFile(fset, path, nil, parser.AllErrors)
	if err != nil {
		log.Println(err)
		return
	}
	// 打印语法树
	ast.Print(fset, f)
}
