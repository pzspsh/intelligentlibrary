/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 16:16:49
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.OpenFile("filepath/hello.go", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("//你好Golang" + strconv.Itoa(i) + "\n")
	}
	writer.Flush()
}
