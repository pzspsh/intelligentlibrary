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
