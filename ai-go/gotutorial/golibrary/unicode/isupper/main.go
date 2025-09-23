package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Printf("%t\n", unicode.IsUpper('A'))
	fmt.Printf("%t\n", unicode.IsUpper('a'))
}
