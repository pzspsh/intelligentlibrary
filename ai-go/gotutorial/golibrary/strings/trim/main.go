package main

import (
	"fmt"
	"strings"
)

func main() { // 去除首尾的指定字符串
	str := " \t\n Hello, Gophers \n\t\r\n"
	fmt.Println(strings.Trim(str, " "))
	fmt.Println(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡"))
	fmt.Println(" hello world")
	fmt.Println(strings.Trim(" hello world", " "))
}
