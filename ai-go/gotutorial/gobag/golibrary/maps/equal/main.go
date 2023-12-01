/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 13:49:25
*/
package main

import (
	"fmt"
	"maps"
)

func main() {
	m := map[string]string{"foo": "bar", "foo2": "bar2"}
	m1 := map[string]string{"foo": "bar2", "foo3": "bar3"}
	m2 := map[string]string{"foo": "bar", "foo2": "bar2"}
	b := maps.Equal(m, m1)
	fmt.Println(b) // false
	b = maps.Equal(m, m2)
	fmt.Println(b) // true
}
