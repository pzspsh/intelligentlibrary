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
	// fmt.Println(strings.Contains("seafood", "foo"))
	// fmt.Println(strings.Contains("seafood", "bar"))
	// fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains(strings.ToLower("RocketMQ <= 5.1.0 - Remote Code Execution"), strings.ToLower("RocketMq Remote")))
}
