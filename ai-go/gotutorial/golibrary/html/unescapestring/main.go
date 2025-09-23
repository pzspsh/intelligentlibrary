package main

import (
	"fmt"
	"html"
)

func main() {
	const s = `&quot;Fran &amp; Freddie&#39;s Diner&quot; &lt;tasty@example.com&gt;`
	fmt.Println(html.UnescapeString(s))
}
