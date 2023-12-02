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
	re := regexp.MustCompile(`(gopher){2}`)
	fmt.Println(re.MatchString("gopher"))
	fmt.Println(re.MatchString("gophergopher"))
	fmt.Println(re.MatchString("gophergophergopher"))
}
