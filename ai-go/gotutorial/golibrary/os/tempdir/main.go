package main

import (
	"fmt"
	"os"
)

func main() {
	tmpDir := os.TempDir()
	fmt.Println("Temporary directory:", tmpDir)
}
