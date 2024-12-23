/*
@File   : main.go
@Author : pan
@Time   : 2023-11-27 18:09:05
*/
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func main() {
	s := strings.NewReader("ABCEFGHIJKLMN")
	br := bufio.NewReader(s)
	b := bytes.NewBuffer(make([]byte, 0))

	br.WriteTo(b)
	fmt.Printf("%s\n", b)
}
