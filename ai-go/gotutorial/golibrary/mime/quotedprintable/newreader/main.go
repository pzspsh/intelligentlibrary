/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 15:24:18
*/
package main

import (
	"fmt"
	"io"
	"mime/quotedprintable"
	"strings"
)

func main() {
	for _, s := range []string{
		`=48=65=6C=6C=6F=2C=20=47=6F=70=68=65=72=73=21`,
		`invalid escape: <b style="font-size: 200%">hello</b>`,
		"Hello, Gophers! This symbol will be unescaped: =3D and this will be written in =\r\none line.",
	} {
		b, err := io.ReadAll(quotedprintable.NewReader(strings.NewReader(s)))
		fmt.Printf("%s %v\n", b, err)
	}
}
