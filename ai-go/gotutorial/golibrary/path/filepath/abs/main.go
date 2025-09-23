package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	val, _ := filepath.Abs("./")
	fmt.Println(val)
}
