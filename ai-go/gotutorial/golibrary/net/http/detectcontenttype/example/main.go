/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 16:49:17
*/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	data := []byte{0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46, 0x00, 0x01}
	contentType := http.DetectContentType(data)
	fmt.Println(contentType)
}
