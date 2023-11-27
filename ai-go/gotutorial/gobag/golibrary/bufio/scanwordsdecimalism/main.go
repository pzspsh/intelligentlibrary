/*
@File   : main.go
@Author : pan
@Time   : 2023-11-24 17:21:46
*/
package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// 使用带有自定义拆分函数(通过包装ScanWords构建)的扫描器来验证32位十进制输入。
func main() {
	// An artificial input source.
	const input = "1234 5678 1234567901234567890"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// Create a custom split function by wrapping the existing ScanWords function.
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		return
	}
	// Set the split function for the scanning operation.
	scanner.Split(split)
	// Validate the input
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
}
