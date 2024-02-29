/*
@File   : main.go
@Author : pan
@Time   : 2024-02-29 14:59:26
*/
package main

import (
	"fmt"
	"regexp"
)

func ExtractLetter(letter string) string {
	reg := regexp.MustCompile("[\u4e00-\u9fa50-9a-zA-Z- ]*")
	letters := reg.FindString(letter)
	return letters
}

func main() {
	letter := ExtractLetter("315soft-FileSystem 2.3.4")
	fmt.Println(letter)
}
