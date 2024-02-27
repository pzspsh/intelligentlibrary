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
	mediatype := "text/html"
	params := map[string]string{
		"charset": "utf-8",
	}

	result := mime.FormatMediaType(mediatype, params)

	fmt.Println("result:", result)
}
