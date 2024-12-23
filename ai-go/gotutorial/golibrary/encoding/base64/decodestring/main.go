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
	str := "c29tZSBkYXRhIHdpdGggACBhbmQg77u/"
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%q\n", data)
}
