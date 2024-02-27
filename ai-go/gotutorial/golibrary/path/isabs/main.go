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
	fmt.Println(path.IsAbs("/dev/null"))
}
