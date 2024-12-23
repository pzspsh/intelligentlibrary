/*
@File   : main.go
@Author : pan
@Time   : 2023-11-27 18:01:12
*/
package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := strings.NewReader("abcdefg")
	br := bufio.NewReader(s)

	fmt.Println(br.Buffered())

	br.Peek(1)
	fmt.Println(br.Buffered())
}
