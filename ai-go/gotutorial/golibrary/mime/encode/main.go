/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 15:19:18
*/
package main

import (
	"fmt"
	"mime"
)

func main() {
	fmt.Println(mime.QEncoding.Encode("utf-8", "¡Hola, señor!"))
	fmt.Println(mime.QEncoding.Encode("utf-8", "Hello!"))
	fmt.Println(mime.BEncoding.Encode("UTF-8", "¡Hola, señor!"))
	fmt.Println(mime.QEncoding.Encode("ISO-8859-1", "Caf\xE9"))
}
