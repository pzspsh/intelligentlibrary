package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.Contains([]byte{1, 2, 3}, []byte{1}))    // true
	fmt.Println(bytes.Contains([]byte{1, 2, 3}, []byte{1, 3})) // false
}
