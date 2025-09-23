package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	src := []byte("Hello")
	encodedStr := hex.EncodeToString(src)

	fmt.Printf("%s\n", encodedStr)

}
