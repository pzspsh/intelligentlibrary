/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 11:53:22
*/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(string(bytes.TrimSpace([]byte(" hello my bro ")))) // hello my bro
}
