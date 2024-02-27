/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 18:45:19
*/
package main

import (
	"log"
	"os"
)

/* 递归创建目录 */
func main() {
	err := os.MkdirAll("test/subdir", 0750)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("test/subdir/testfile.txt", []byte("Hello, Gophers!"), 0660)
	if err != nil {
		log.Fatal(err)
	}
}

/*

创建一个新目录，该目录是利用路径（包括绝对路径和相对路径）进行创建的，如果需要创建对应的父目录，也一起进行创建，
如果已经有了该目录，则不进行新的创建，当创建一个已经存在的目录时，不会报错.
*/
