/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 16:45:30
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getgroups())                //获取调用者属于的组  [4 24 27 30 46 108 124 1000]
	fmt.Println(os.Getgid())                   //获取调用者当前所在的组　1000
	fmt.Println(os.Chown("tmp.txt", 1000, 46)) //更改文件所在的组
}
