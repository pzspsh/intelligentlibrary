/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 13:48:56
*/
package main

import (
	"fmt"
	"maps"
)

func main() {
	m := map[string]string{"foo": "bar", "foo2": "bar2"}
	m1 := map[string]string{"foo": "bar2", "foo3": "bar3"}
	maps.Copy(m1, m)
	fmt.Println(m1) // map[foo:bar foo2:bar2 foo3:bar3]
}
