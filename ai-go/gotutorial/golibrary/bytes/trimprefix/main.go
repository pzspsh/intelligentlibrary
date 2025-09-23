package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b = []byte("Goodbye,, world!")
	b = bytes.TrimPrefix(b, []byte("Goodbye,"))
	b = bytes.TrimPrefix(b, []byte("See ya,"))
	fmt.Printf("Hello%s", b)                                                // hello, world!
	fmt.Println(string(bytes.TrimPrefix([]byte("hi hi go"), []byte("hi")))) //  hi go
}
