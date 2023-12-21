/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 18:45:19
*/
package main

import (
	"fmt"
	"log"
	"os"
)

/* 创建目录 */
func main() {
	if err := os.Mkdir("infodir", os.ModeDir); err != nil {
		fmt.Println("error")
	} else {
		fmt.Println("success")
	}

	err := os.Mkdir("testdir", 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	err = os.WriteFile("testdir/testfile.txt", []byte("Hello, Gophers!"), 0660)
	if err != nil {
		log.Fatal(err)
	}
}
