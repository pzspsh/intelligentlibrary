package main

import (
	"fmt"
	"os"
)

// ioutil.ReadFile读取整个文件
func main() {
	content, err := os.ReadFile("./main.go")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
}
