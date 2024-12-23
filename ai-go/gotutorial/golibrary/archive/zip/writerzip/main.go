/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 10:26:25
*/
package main

import (
	"archive/zip"
	"fmt"
	"os"
)

func main() {
	const dir = "filepath/test001/"
	//获取源文件列表
	f, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}
	fzip, _ := os.Create("filepath/test001/test00001.zip")
	w := zip.NewWriter(fzip)
	defer w.Close()
	for _, file := range f {
		fw, _ := w.Create(file.Name())
		filecontent, err := os.ReadFile(dir + file.Name())
		if err != nil {
			fmt.Println(err)
		}
		n, err := fw.Write(filecontent)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(n)
	}
}
