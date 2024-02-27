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
	const s = `&quot;Fran &amp; Freddie&#39;s Diner&quot; &lt;tasty@example.com&gt;`
	fmt.Println(html.UnescapeString(s))
}
