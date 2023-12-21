/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 15:57:36
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	fi, _ := os.Lstat("filepath/main.go")

	fmt.Println(fi.Size())
}
