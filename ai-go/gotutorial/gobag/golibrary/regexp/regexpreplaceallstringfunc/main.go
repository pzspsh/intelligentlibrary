/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 21:28:45
*/
package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	re := regexp.MustCompile(`[^aeiou]`)
	fmt.Println(re.ReplaceAllStringFunc("seafood fool", strings.ToUpper))
}
