/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 10:14:35
*/
package main

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Create("filepath/test10.tar") //创建一个tar文件
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// NewWriter
	tw := tar.NewWriter(f)
	defer tw.Close()

	fileinfo, err := os.Stat("filepath/test01.sh") //获取文件相关信息
	if err != nil {
		fmt.Println(err)
	}
	hdr, err := tar.FileInfoHeader(fileinfo, "")
	if err != nil {
		fmt.Println(err)
	}

	// WriterHeader
	err = tw.WriteHeader(hdr) //写入头文件信息
	if err != nil {
		fmt.Println(err)
		// return
	}

	f1, err := os.Open("filepath/test01.sh")
	if err != nil {
		fmt.Println(err)
		return
	}
	m, err := io.Copy(tw, f1) //将文件test01.sh中信息写入压缩包中
	if err != nil {
		fmt.Println(err)
		// return
	}

	fileinfo2, err := os.Stat("filepath/test02.sh") //获取文件相关信息
	if err != nil {
		fmt.Println(err)
	}
	hdr2, err := tar.FileInfoHeader(fileinfo2, "")
	if err != nil {
		fmt.Println(err)
	}

	// WriterHeader
	err = tw.WriteHeader(hdr2) //写入头文件信息
	if err != nil {
		fmt.Println(err)
		// return
	}

	f2, err := os.Open("filepath/test02.sh")
	if err != nil {
		fmt.Println(err)
		return
	}
	m2, err := io.Copy(tw, f2) //将文件test02.sh中信息写入压缩包中
	if err != nil {
		fmt.Println(err)
		// return
	}

	fmt.Println(m)
	fmt.Println(m2)
}
