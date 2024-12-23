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
	fmt.Println("On Unix:")
	fmt.Println(filepath.Join("a", "b", "c"))
	fmt.Println(filepath.Join("a", "b/c"))
	fmt.Println(filepath.Join("a/b", "c"))
	fmt.Println(filepath.Join("a/b", "/c"))

	fmt.Println(filepath.Join("a/b", "../../../xyz"))
}
