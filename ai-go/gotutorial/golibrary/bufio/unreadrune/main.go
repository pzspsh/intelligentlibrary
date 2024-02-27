/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 09:38:30
*/
package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := strings.NewReader("你好，世界！")
	br := bufio.NewReader(s)

	c, size, _ := br.ReadRune()
	fmt.Printf("%c %v\n", c, size)

	c, size, _ = br.ReadRune()
	fmt.Printf("%c %v\n", c, size)

	br.UnreadRune()
	c, size, _ = br.ReadRune()
	fmt.Printf("%c %v\n", c, size)
}
