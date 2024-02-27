/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 23:08:40
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("fail", "I"))
	fmt.Println(strings.ContainsAny("ure", "ui"))
	// fmt.Println(strings.ContainsAny("failure", "ui"))
	// fmt.Println(strings.ContainsAny("foo", ""))
	// fmt.Println(strings.ContainsAny("", ""))
}
