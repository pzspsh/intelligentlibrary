/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 15:33:19
*/
package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// 压缩
	tarFile := "filepath/demo.tar.gz"
	dest := "filepath/test001/"

	srcFile, err := os.Open(tarFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer srcFile.Close()
	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
		filename := dest + hdr.Name
		err = os.MkdirAll(string([]rune(filename)[0:strings.LastIndex(filename, "/")]), 0755)
		if err != nil {
			fmt.Println(err)
			return
		}
		file, _ := os.Create(filename)
		io.Copy(file, tr)
	}
}
