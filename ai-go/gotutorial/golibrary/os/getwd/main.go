package main

import (
	"fmt"
	"os"
)

/* 获取当前工作目录 */
func main() {
	path, _ := os.Getwd()
	fmt.Println(path)
}
