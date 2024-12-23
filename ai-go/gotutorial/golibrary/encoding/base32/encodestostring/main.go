/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 14:46:13
*/
package main

import (
	"encoding/base32"
	"fmt"
)

func main() {
	data := []byte("any + old & data")
	str := base32.StdEncoding.EncodeToString(data)
	fmt.Println(str)
}
