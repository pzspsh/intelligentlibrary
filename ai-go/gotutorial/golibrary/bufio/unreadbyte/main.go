package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := strings.NewReader("ABCDEFG")
	br := bufio.NewReader(s)

	c, _ := br.ReadByte()
	fmt.Printf("%c\n", c)

	c, _ = br.ReadByte()
	fmt.Printf("%c\n", c)

	br.UnreadByte()
	c, _ = br.ReadByte()
	fmt.Printf("%c\n", c)
}
