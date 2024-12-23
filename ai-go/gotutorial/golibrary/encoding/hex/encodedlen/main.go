/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 15:34:00
*/
package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	srccode := []byte("Hello world!")
	dstEncode := make([]byte, hex.EncodedLen(len(srccode)))
	hex.Encode(dstEncode, srccode)
	fmt.Printf("%x\n", dstEncode)
	data, _ := hex.DecodeString(string(dstEncode))
	fmt.Printf("%x\n", data)
}
