package main

import (
	"fmt"
	"unicode"
)

func main() {
	const ucG = 'g'
	fmt.Printf("%#U\n", unicode.ToTitle(ucG))
}
