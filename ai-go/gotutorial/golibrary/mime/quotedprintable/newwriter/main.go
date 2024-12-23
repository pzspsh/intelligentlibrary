/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 15:24:18
*/
package main

import (
	"mime/quotedprintable"
	"os"
)

func main() {
	w := quotedprintable.NewWriter(os.Stdout)
	w.Write([]byte("These symbols will be escaped: = \t"))
	w.Close()

}
