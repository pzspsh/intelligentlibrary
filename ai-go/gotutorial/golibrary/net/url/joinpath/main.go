/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 17:08:12
*/

package main

import (
	"fmt"
	"net/url"
	"path"
)

func joinURLPath(baseURL, newPath string) (string, error) {
	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse base URL: %v", err)
	}
	joinedPath := path.Join(parsedBaseURL.Path, newPath)
	parsedBaseURL.Path = joinedPath
	return parsedBaseURL.String(), nil
}

func main() {
	baseURL := "https://example.com/api/v1"
	newPath := "users"

	joinedURL, err := joinURLPath(baseURL, newPath)
	if err != nil {
		fmt.Printf("Error joining URL paths: %v\n", err)
		return
	}

	fmt.Printf("Joined URL: %s\n", joinedURL)
}
