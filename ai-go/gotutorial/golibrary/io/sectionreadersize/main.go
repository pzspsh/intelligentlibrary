/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 12:39:38
*/
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	s := io.NewSectionReader(r, 5, 17)

	fmt.Println(s.Size())

}
