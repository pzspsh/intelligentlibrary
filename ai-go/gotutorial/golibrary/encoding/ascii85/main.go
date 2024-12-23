/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 14:25:31
*/
package main

import (
	"encoding/ascii85"
	"fmt"
)

func main() {
	dst := make([]byte, 25)
	dst2 := make([]byte, 25)
	ascii85.Encode(dst, []byte("Hello, playground"))
	fmt.Println(dst)
	ascii85.Decode(dst2, dst, false)
	fmt.Println(string(dst2))
}
