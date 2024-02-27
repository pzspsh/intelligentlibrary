/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 10:11:04
*/
package main

import (
	"archive/tar"
	"fmt"
	"os"
)

func main() {
	fileinfo, err := os.Stat("path/test.docx")
	if err != nil {
		fmt.Println(err)
	}
	h, err := tar.FileInfoHeader(fileinfo, "")
	h.Linkname = "linkname"
	h.Gname = "gname"
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(h.AccessTime, h.ChangeTime, h.Devmajor, h.Devminor, h.Gid, h.Gname, h.Linkname, h.ModTime, h.Mode, h.Name, h.Size, h.Typeflag, h.Uid, h.Uname, h.Xattrs)

	// FileInfo
	fileinfo2 := h.FileInfo()
	fmt.Println(fileinfo2.Name())
}
