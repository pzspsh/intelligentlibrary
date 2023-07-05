package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(simplifyPath1("/home/"))
}

func simplifyPath1(path string) string {
	return filepath.Clean(path)
}
