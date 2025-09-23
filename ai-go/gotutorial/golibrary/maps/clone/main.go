package main

import (
	"fmt"
	"maps"
)

func main() {
	m := map[string]string{"foo": "bar"}
	m1 := maps.Clone(m)
	fmt.Println(m1) // map[foo:bar]
}
