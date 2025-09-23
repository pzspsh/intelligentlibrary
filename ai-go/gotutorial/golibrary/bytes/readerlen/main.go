package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.NewReader([]byte("Hi!")).Len())    // 3
	fmt.Println(bytes.NewReader([]byte("こんにちは!")).Len()) // 16
}
