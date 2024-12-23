/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 15:32:59
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.CreateTemp("", "sample")
	check(err)

	fmt.Println("Temp file name:", f.Name())
	// 操作系统会在程序结束后一段时间自动清理临时文件，不过最好还是自己做清理
	defer os.Remove(f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
