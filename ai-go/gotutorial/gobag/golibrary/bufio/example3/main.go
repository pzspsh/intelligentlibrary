/*
@File   : main.go
@Author : pan
@Time   : 2023-11-27 09:28:00
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(strings.NewReader("gopher"))
	for scanner.Scan() {
		fmt.Println(len(scanner.Bytes()) == 6)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "shouldn't see an error scanning a string")
	}
}
