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
