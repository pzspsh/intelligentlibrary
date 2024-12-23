/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 14:14:08
*/
package main

import (
	_ "embed"
	"fmt"
)

/*
go:embed version.txt info.txt
*/

//go:embed info.txt
var version string

func main() {
	fmt.Printf("version %q\n", version)
}
