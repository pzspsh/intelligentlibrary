/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 10:25:26
*/
package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func main() {
	const File = "filepath/test001.zip"
	const dir = "filepath/test001/"
	os.Mkdir(dir, 0777) //创建一个目录

	or, err := zip.OpenReader(File) //读取zip文件
	if err != nil {
		fmt.Println(err)
	}
	defer or.Close()
	for _, file := range or.File {
		rc, err := file.Open()
		if err != nil {
			fmt.Println(err)
		}

		f, err := os.Create(dir + file.Name)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		n, err := io.Copy(f, rc)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(n)
	}
}
