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
	fmt.Println(regexp.QuoteMeta(`Escaping symbols like: .+*?()|[]{}^$`))
}
