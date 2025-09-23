package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	data := "abc123"
	reader := bufio.NewReader(strings.NewReader(data))
	for {
		c, err := reader.ReadByte()
		if err != nil {
			break
		}
		fmt.Print(string(c))
	}
}
