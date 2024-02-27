/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 16:18:02
*/
package main

import (
	"fmt"
	"io"
	"os"
)

// 编写一个函数，接受两个文件路径 srcFileName 和 dstFileName
func CopyFile(srcFileName string, dstFileName string) (err error) {
	sFile, err1 := os.Open(srcFileName)
	if err1 != nil {
		return err1
	}
	defer sFile.Close()
	dFile, err2 := os.OpenFile(dstFileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err2 != nil {
		return err2
	}
	defer dFile.Close()
	var tempSlice = make([]byte, 50)
	for {
		//读取数据
		n, err := sFile.Read(tempSlice)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		//写入数据
		if _, err := dFile.Write(tempSlice[:n]); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	srcFile := "filepath/hello.go"
	dstFile := "filepath/world.go"
	err := CopyFile(srcFile, dstFile)
	if err == nil {
		fmt.Println("拷贝完成")
	} else {
		fmt.Printf("拷贝错误 err = %v", err)
	}
}
