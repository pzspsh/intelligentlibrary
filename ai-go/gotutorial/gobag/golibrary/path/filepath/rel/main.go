/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 20:13:53
*/
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	paths := []string{
		"/a/b/c",
		"/b/c",
		"./b/c",
	}
	base := "/a"

	fmt.Println("On Unix:")
	for _, p := range paths {
		rel, err := filepath.Rel(base, p)
		fmt.Printf("%q: %q %v\n", p, rel, err)
	}
}
