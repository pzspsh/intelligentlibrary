/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 14:53:51
*/
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("any + old & data")
	str := base64.StdEncoding.EncodeToString(data)
	fmt.Println(str)
}
