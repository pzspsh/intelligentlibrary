package main

import (
	"fmt"
	"strconv"
)

func main() {
	// This string literal contains a tab character.
	s := strconv.Quote(`"Fran & Freddie's Diner	☺"`)
	fmt.Println(s)
}
