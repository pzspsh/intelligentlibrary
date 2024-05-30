/*
@File   : main.go
@Author : pan
@Time   : 2024-05-30 16:10:44
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	tmpDir := os.TempDir()
	fmt.Println("Temporary directory:", tmpDir)
}
