package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%s", bytes.ToUpper([]byte("Gopher"))) // GOPHER
	fmt.Println(string(bytes.ToUpper([]byte("Aaa")))) // AAA
}
