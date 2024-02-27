/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 20:13:53
*/
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println("On Unix:", filepath.SplitList("/a/b/c:/usr/bin"))
}
