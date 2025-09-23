package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(string(bytes.Title([]byte("AAA")))) // AAA
	fmt.Println(string(bytes.Title([]byte("aaa")))) // Aaa
}
