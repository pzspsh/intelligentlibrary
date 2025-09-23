package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := strings.NewReader("ABCEFG")
	str := strings.NewReader("123455")
	br := bufio.NewReader(s)
	b, _ := br.ReadString('\n')
	fmt.Println(b)
	br.Reset(str)
	b, _ = br.ReadString('\n')
	fmt.Println(b)
}
