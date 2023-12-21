/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 16:46:43
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.IsPathSeparator('/')) //true
	fmt.Println(os.IsPathSeparator('|')) //false
}
