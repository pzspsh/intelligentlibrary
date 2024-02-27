/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 10:12:52
*/
package main

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
)

func main() {

	f, err := os.Open("path/test.tar")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	// NewReader
	r := tar.NewReader(f)
	// Next
	for hdr, err := r.Next(); err != io.EOF; hdr, err = r.Next() {
		if err != nil {
			fmt.Println(err)
			return
		}
		fileinfo := hdr.FileInfo()
		fmt.Println(fileinfo.Name())
		f, err := os.Create("path/" + fileinfo.Name())
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		_, err = io.Copy(f, r)
		if err != nil {
			fmt.Println(err)
		}
	}
}
