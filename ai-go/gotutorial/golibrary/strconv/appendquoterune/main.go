/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 22:32:49
*/
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
