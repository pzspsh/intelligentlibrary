/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 15:34:03
*/
package main

import (
	"embed"
)

// 使用//go:embed指定从文件读取内容到变量fileString、fileByte
//
//go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileByte []byte

//使用go:embed也可以打包多个文件到embed.FS类型变量，方便后续操作

//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {

	print(fileString)
	print(string(fileByte))

	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}
