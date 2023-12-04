/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 17:09:05
*/
package main

import (
	"fmt"
	"net/url"
	"path"
)

func createURL(baseURL, user string) (string, error) {
	// Parse the base URL
	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse base URL: %v", err)
	}

	// Escape the user string
	escapedUser := url.PathEscape(user)

	// Join the base URL path with the escaped user string
	joinedPath := path.Join(parsedBaseURL.Path, escapedUser)

	// Set the new path
	parsedBaseURL.Path = joinedPath

	// Convert the parsed URL back to a string
	return parsedBaseURL.String(), nil
}

func main() {
	baseURL := "https://example.com/users"
	user := "John Doe !@#$%^&*()"

	// Create the URL
	newURL, err := createURL(baseURL, user)
	if err != nil {
		fmt.Printf("Error creating URL: %v\n", err)
		return
	}

	fmt.Printf("Generated URL: %s\n", newURL)
}
