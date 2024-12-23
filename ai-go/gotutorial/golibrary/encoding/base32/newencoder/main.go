/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 14:46:13
*/
package main

import (
	"encoding/base32"
	"os"
)

func main() {
	input := []byte("foo\x00bar")
	encoder := base32.NewEncoder(base32.StdEncoding, os.Stdout)
	encoder.Write(input)
	// Must close the encoder when finished to flush any partial blocks.
	// If you comment out the following line, the last partial block "r"
	// won't be encoded.
	encoder.Close()
}
