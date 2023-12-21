/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 18:45:19
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	var path string
	if os.IsPathSeparator('\\') {
		path = "\\"
	} else {
		path = "/"
	}
	pwd, _ := os.Getwd()
	err := os.Mkdir(pwd+path+"tmp", os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}
