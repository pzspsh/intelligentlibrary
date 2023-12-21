/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 16:26:46
*/
package main

import (
	"fmt"
	"log"
	"os"
)

/*
功能：
读取并返回目录f里面的文件(或文件夹)的名称列表

如果n>0，Readdirnames返回最多n个名字。

如果n<0，Readdirnames返回目录下所有的文件的名字，用一个切片表示。
*/

func main() {
	curPwd, _ := os.Getwd()
	dirOp, err := os.Open(curPwd)
	if err != nil {
		log.Fatal(err)
	}
	defer dirOp.Close()
	dirList, err := dirOp.Readdirnames(-1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dirList)
}
