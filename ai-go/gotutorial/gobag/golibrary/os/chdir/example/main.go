/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 16:24:21
*/
package main

import (
	"fmt"
	"os"
)

/*
功能：改变工作目录到f，其中f必须为一个目录，否则便会报错。目前不支持windows系统，如果windows系统下可用os.Chdir()函数来改变工作目录。
*/
func main() {
	dir, _ := os.Getwd()
	fmt.Println(dir)
	f, _ := os.Open("tmp.txt")
	err := f.Chdir()
	if err != nil {
		fmt.Println(err) //chdir tmp.txt: not a directory，因为tmp.txt不是目录，所以报错
	}
	f, _ = os.Open("develop")
	err = f.Chdir()
	if err != nil {
		fmt.Println(err)
	}
	dir1, _ := os.Getwd()
	fmt.Println(dir1) //home/work/develop，因为develop是工作目录，所以切换到该目录
}
