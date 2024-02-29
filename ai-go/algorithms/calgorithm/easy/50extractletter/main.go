/*
@File   : main.go
@Author : pan
@Time   : 2024-02-29 14:59:26
*/
package main

import (
	"fmt"
	"regexp"
	"strings"
)

func ExtractLetter(letter string) string {
	reg := regexp.MustCompile(`[a-zA-Z-]`)
	letters := reg.FindAllString(letter, -1)
	letterstr := strings.Join(letters, "")
	return letterstr
}

func main() {
	letter := ExtractLetter("Apache-Coyote/1.1")
	fmt.Println(letter)
}
