/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 12:33:19
*/
package main

import (
	"fmt"
	"html"
)

func main() {
	const s = `"Fran & Freddie's Diner" <tasty@example.com>`
	fmt.Println(html.EscapeString(s))
}
