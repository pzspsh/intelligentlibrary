package main

import (
	"fmt"
	"os"
)

/* 修改工作目录 */
func main() {
	path1, _ := os.Getwd()
	fmt.Println(path1)
	os.Chdir("./../")
	path, _ := os.Getwd()
	fmt.Println(path)
}
