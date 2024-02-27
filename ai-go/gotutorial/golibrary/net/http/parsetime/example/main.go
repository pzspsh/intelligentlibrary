/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 16:40:58
*/
package main

import (
	"fmt"

	"net/http"
)

func main() {
	httpTimeString := "Sun, 06 Nov 1994 08:49:37 GMT"
	t, err := http.ParseTime(httpTimeString)
	if err == nil {
		fmt.Printf("Parsed time: %s\n", t)

	} else {
		fmt.Println("Error parsing HTTP time:", err)
	}
}
