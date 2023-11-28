/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 10:23:27
*/
package main

import (
	"archive/zip"
	"fmt"
	"os"
	"time"
)

func main() {
	fileinfo, err := os.Stat("filepath/test.docx")
	if err != nil {
		fmt.Println(err)
	}
	h, err := zip.FileInfoHeader(fileinfo)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(h.Name, h.Comment, h.CreatorVersion)

	// FileInfo
	fileinfo2 := h.FileInfo()
	fmt.Println(fileinfo2.Name())

	// Mode
	fmt.Println(h.Mode())
	// SetMode
	h.SetMode(0755)
	// h.SetMode(755)
	fmt.Println(h.Mode())

	// Deprecated: Use Modified instead.
	//modTime := h.ModTime()
	//fmt.Println(modTime)
	//
	// Deprecated: Use Modified instead.
	//h.SetModTime(time.Now())
	//fmt.Println(h.ModTime())

	fmt.Println(h.Modified)
	h.Modified = time.Now()
	fmt.Println(h.Modified)
}
