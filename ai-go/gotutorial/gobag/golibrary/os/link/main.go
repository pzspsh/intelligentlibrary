/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 16:21:29
*/
package main

import (
	"fmt"
	"log"
	"os"
)

/*
func Link(oldname, newname string) error

功能：创建一个从oldname指向newname的硬连接，对一个进行操作，则另外一个也会被修改。

注意：newname必须是不存在的，负责会返回error（Cannot create a file when that file already exists.）
*/
func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dir1 := dir + "/filepath/a.go"
	dir2 := dir + "/filepath/b.go"
	err = os.Link(dir1, dir2) //dir1必须存在，dir2必须是不存在的，负责会报错
	if err != nil {
		log.Fatal("create link error ", err)
	}

	fil1, err := os.OpenFile(dir1, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0555)
	if err != nil {
		log.Fatal("open file1 error  ", err)
	}
	fil2, err := os.OpenFile(dir2, os.O_RDWR, 0555)
	if err != nil {
		log.Fatal("open file2 error  ", err)
	}

	fil1.WriteString(`
package main
import "fmt"

func main(){
	 fmt.Println("hello link")
}
`)

	bytes := make([]byte, 128)
	n, err := fil2.Read(bytes)

	if err != nil {
		log.Fatal("read file2 error  ", err)
	}

	fmt.Println(string(bytes[:n]))
}
