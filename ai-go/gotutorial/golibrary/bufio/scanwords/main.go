/*
@File   : main.go
@Author : pan
@Time   : 2023-11-27 09:25:18
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 使用Scanner实现简单的单词计数实用程序，方法是将输入扫描为以空格分隔的令牌序列。
func main() {
	// An artificial input source.
	const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Count the words.
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("%d\n", count)
}
