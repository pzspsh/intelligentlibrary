package main

import (
	"fmt"
	"html"
)

func main() {
	const s = `"Fran & Freddie's Diner" <tasty@example.com>`
	fmt.Println(html.EscapeString(s))
}
