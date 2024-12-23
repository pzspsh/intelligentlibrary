/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 20:34:49
*/
package main

import (
	"fmt"
	"plugin"
)

var V int

func F() {
	fmt.Printf("Hello, number %d\n", V)
}

// 把以上F()打包成.so文件

func main() {
	p, err := plugin.Open("plugin_name.so")
	if err != nil {
		panic(err)
	}
	v, err := p.Lookup("V")
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("F")
	if err != nil {
		panic(err)
	}
	*v.(*int) = 7
	f.(func())() // prints "Hello, number 7"
}
