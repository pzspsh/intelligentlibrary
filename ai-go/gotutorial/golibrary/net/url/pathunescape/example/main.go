/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 17:09:05
*/
package main

import (
	"fmt"
	"net/url"
)

func decodeURLPath(encodedPath string) (string, error) {
	// Decode the encoded path using PathUnescape
	decodedPath, err := url.PathUnescape(encodedPath)

	if err != nil {
		return "", fmt.Errorf("failed to decode path: %v", err)
	}

	return decodedPath, nil
}

func main() {
	encodedPath := "/users/John%20Doe%20%21%40%23%24%25%5E%26%2A%28%29"

	// Decode the URL path
	decodedPath, err := decodeURLPath(encodedPath)

	if err != nil {
		fmt.Printf("Error decoding URL path: %v\n", err)
		return
	}

	fmt.Printf("Decoded path: %s\n", decodedPath)
}
