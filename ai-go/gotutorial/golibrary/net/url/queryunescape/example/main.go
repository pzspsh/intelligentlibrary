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

func decodeQueryParameter(encodedParam string) (string, error) {
	// Decode the encoded parameter using QueryUnescape
	decodedParam, err := url.QueryUnescape(encodedParam)

	if err != nil {
		return "", fmt.Errorf("failed to decode query parameter: %v", err)
	}

	return decodedParam, nil
}

func main() {
	encodedParam := "Go+language+%26+net%2Furl+package"

	// Decode the query parameter
	decodedParam, err := decodeQueryParameter(encodedParam)

	if err != nil {
		fmt.Printf("Error decoding query parameter: %v\n", err)
		return
	}

	fmt.Printf("Decoded query parameter: %s\n", decodedParam)
}
