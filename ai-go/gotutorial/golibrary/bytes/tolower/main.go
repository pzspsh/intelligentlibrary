package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%s", bytes.ToLower([]byte("Gopher"))) // gopher
	fmt.Println(string(bytes.ToLower([]byte("aAA")))) // aaa
}
