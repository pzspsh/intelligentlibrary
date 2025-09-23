package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := []byte("rune:")
	b = strconv.AppendQuoteRune(b, '☺')
	fmt.Println(string(b))
}
