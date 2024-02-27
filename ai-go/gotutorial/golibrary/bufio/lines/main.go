/*
@File   : main.go
@Author : pan
@Time   : 2023-11-27 09:22:57
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

// 扫描器的最简单用途，将标准输入读取为一组行。
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
