/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 14:11:51
*/
package main

import (
	_ "embed"
	"fmt"
)

//go:embed version.txt
var version string

func main() {
	fmt.Printf("version: %q\n", version)
}
