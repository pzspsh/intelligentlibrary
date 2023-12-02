/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 21:28:45
*/
package main

import (
	"fmt"
	"regexp"
)

func main() {
	content := []byte("London")
	re := regexp.MustCompile(`o.`)
	fmt.Println(re.FindAllIndex(content, 1))
	fmt.Println(re.FindAllIndex(content, -1))
}
