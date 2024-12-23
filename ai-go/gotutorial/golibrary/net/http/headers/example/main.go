/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 16:50:19
*/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	header := http.Header{}

	header.Add("content-type", "application/json")
	header.Add("x-custom-header", "custom-value")

	for key, value := range header {
		fmt.Printf("Key: %s, Value: %s\n", http.CanonicalHeaderKey(key), value)
	}
}
