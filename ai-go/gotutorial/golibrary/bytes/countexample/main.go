/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 12:26:20
*/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.Count([]byte("cheese"), []byte("e")))
	fmt.Println(bytes.Count([]byte("five"), []byte(""))) // before & after each rune
}
