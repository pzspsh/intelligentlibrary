/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 15:30:33
*/
package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	// 将不同级别目录名拼接成文件的完整路径
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))
	// 从完整路径获取目录名和文件名
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))
	// 是否是绝对路径
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"
	// 获取文件扩展名
	ext := filepath.Ext(filename)
	fmt.Println(ext)
	// 获取去除扩展名后的文件名
	fmt.Println(strings.TrimSuffix(filename, ext))
	// 获取两个路径之间的相对路径
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
