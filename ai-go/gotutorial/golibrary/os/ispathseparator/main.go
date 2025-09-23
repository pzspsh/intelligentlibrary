package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.IsPathSeparator('/')) //true
	fmt.Println(os.IsPathSeparator('|')) //false
}
