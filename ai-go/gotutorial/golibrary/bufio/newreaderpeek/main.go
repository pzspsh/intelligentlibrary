/*
@File   : main.go
@Author : pan
@Time   : 2023-11-27 18:02:06
*/
package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	br := bufio.NewReader(s)

	b, _ := br.Peek(5)
	fmt.Printf("%s\n", b)

	b[0] = 'a'
	b, _ = br.Peek(5)
	fmt.Printf("%s\n", b)
}
