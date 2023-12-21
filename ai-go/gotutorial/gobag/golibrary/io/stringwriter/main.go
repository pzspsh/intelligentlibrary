/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 17:09:07
*/
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	var builder strings.Builder
	writer := io.StringWriter(&builder)
	writer.WriteString("Hello, ")
	writer.WriteString("World!")
	result := builder.String()
	fmt.Println(result) // 输出：Hello, World!
}
