/*
@File   : main.go
@Author : pan
@Time   : 2023-12-07 16:50:45
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	gid := os.Getegid()
	fmt.Println(gid)
	dirfs := os.DirFS(`path/gotutorial/gobag/golibrary/os/writefile/main.go`) // DirFS为目录dir下的文件树返回一个文件系统(fs.FS)。
	fmt.Println(dirfs)
}
