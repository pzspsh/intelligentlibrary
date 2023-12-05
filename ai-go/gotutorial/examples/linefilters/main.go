/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 15:30:02
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 从标准输入创建scanner
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// 默认用空格分隔，每次新读取一段
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
