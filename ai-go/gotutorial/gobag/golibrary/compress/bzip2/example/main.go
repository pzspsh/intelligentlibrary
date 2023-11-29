/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 15:31:34
*/
package main

import (
	"compress/bzip2"
	"fmt"
	"io"
	"os"
)

func main() {
	zip2File := "filepath/test.txt.bz2"
	unzip2FileName := "filepath/unzipfile2.txt"
	//打开要解包的文件，zip2File是要解包的 bz2 文件的路径
	fr, err := os.Open(zip2File)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer fr.Close()
	// 创建 bzip2.Reader，准备执行解包操作
	br := bzip2.NewReader(fr)
	//创建空文件，准备写入解压后的数据
	bw, err := os.Create(unzip2FileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer bw.Close()
	// 写入解压后的数据
	_, err = io.Copy(bw, br)
	if err != nil {
		fmt.Println(err.Error())
	}
}
