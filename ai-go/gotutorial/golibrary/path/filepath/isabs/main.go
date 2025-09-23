package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println("On Unix:")
	fmt.Println(filepath.IsAbs("/home/gopher"))
	fmt.Println(filepath.IsAbs(".bashrc"))
	fmt.Println(filepath.IsAbs(".."))
	fmt.Println(filepath.IsAbs("."))
	fmt.Println(filepath.IsAbs("/"))
	fmt.Println(filepath.IsAbs(""))
}
