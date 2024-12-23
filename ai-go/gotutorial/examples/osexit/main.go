/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 15:42:57
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("!")
	os.Exit(3)
}
