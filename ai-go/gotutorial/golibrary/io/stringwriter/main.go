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
