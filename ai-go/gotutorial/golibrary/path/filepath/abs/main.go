/*
@File   : main.go
@Author : pan
@Time   : 2024-05-30 16:02:09
*/
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	val, _ := filepath.Abs("./")
	fmt.Println(val)
}
