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

func createURLWithQuery(baseURL, key, value string) (string, error) {
	// Parse the base URL
	parsedBaseURL, err := url.Parse(baseURL)

	if err != nil {
		return "", fmt.Errorf("failed to parse base URL: %v", err)
	}

	// Escape the key and value
	escapedKey := url.QueryEscape(key)
	escapedValue := url.QueryEscape(value)

	// Create a new query and add the escaped key-value pair
	query := url.Values{}
	query.Add(escapedKey, escapedValue)

	// Set the new query to the parsed base URL
	parsedBaseURL.RawQuery = query.Encode()

	// Convert the parsed URL back to a string
	return parsedBaseURL.String(), nil
}

func main() {
	baseURL := "https://example.com/search"
	key := "query"
	value := "Go language & net/url package"

	// Create the URL with query parameters
	newURL, err := createURLWithQuery(baseURL, key, value)

	if err != nil {
		fmt.Printf("Error creating URL: %v\n", err)
		return
	}

	fmt.Printf("Generated URL: %s\n", newURL)
}
