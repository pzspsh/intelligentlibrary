/*
@File   : fileexists.go
@Author : pan
@Time   : 2023-06-21 11:07:47
*/
package main

import (
	"fmt"
	"os"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func main() {
	result := FileExists("../tests/test.txt")
	fmt.Println(result)
}
