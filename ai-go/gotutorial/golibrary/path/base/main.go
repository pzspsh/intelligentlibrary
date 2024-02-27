/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 19:15:21
*/
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Base("/a/b"))
	fmt.Println(path.Base("/"))
	fmt.Println(path.Base(""))
}
