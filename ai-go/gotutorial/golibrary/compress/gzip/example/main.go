/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 15:34:47
*/
package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	// 压缩
	dest := "filepath/demo.tar.gz"
	d, _ := os.Create(dest)

	defer d.Close()
	gw := gzip.NewWriter(d)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	file_dirs := []string{"filepath/test01.sh", "filepath/test02.sh"}

	var files []*os.File
	for _, dir := range file_dirs {
		file, _ := os.Open(dir)
		files = append(files, file)
	}

	for _, file := range files {
		info, err := file.Stat()
		if err != nil {
			fmt.Println(err)
			return
		}
		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = tw.WriteHeader(header)
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = io.Copy(tw, file)
		file.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
