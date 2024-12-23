/*
@File   : main.go
@Author : pan
@Time   : 2023-11-27 10:29:49
*/
package main

import (
	"fmt"
	"log"
	"os"
)

var (
	fileInfo os.FileInfo
	err      error
)

func main() {
	// 如果文件不存在，则返回错误
	fileInfo, err = os.Stat("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
}
