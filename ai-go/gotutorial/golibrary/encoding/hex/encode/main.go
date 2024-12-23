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
	src := []byte("Hello Gopher!")

	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)

	fmt.Printf("%s\n", dst)

}
