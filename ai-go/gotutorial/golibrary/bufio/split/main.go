/*
@File   : main.go
@Author : pan
@Time   : 2023-11-27 09:20:27
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 使用带有自定义split函数的Scanner来解析带有空最终值的逗号分隔列表。EmptyFinalToken
func main() {
	// Comma-separated list; last entry is empty.
	const input = "1,2,3,4,"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// Define a split function that separates on commas.--> SplitFunc
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}
		// There is one final token to be delivered, which may be the empty string.
		// Returning bufio.ErrFinalToken here tells Scan there are no more tokens after this
		// but does not trigger an error to be returned from Scan itself.
		return 0, data, bufio.ErrFinalToken
	}
	scanner.Split(onComma)
	// Scan.
	for scanner.Scan() {
		fmt.Printf("%q ", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
