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
	str := "ONXW2ZJAMRQXIYJAO5UXI2BAAAQGC3TEEDX3XPY="
	data, err := base32.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%q\n", data)
}
