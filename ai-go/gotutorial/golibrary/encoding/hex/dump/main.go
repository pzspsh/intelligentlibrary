/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 15:31:52
*/
package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	content := []byte("Go is an open source programming language.")

	fmt.Printf("%s", hex.Dump(content))

}
