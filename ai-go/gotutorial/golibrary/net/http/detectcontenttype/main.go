/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:06:42
*/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	cont1 := http.DetectContentType([]byte{})                                           //text/plain; charset=utf-8
	cont2 := http.DetectContentType([]byte{1, 2, 3})                                    //application/octet-stream
	cont3 := http.DetectContentType([]byte(`<HtMl><bOdY>blah blah blah</body></html>`)) //text/html; charset=utf-8
	cont4 := http.DetectContentType([]byte("\n<?xml!"))                                 //text/xml; charset=utf-8
	cont5 := http.DetectContentType([]byte(`GIF87a`))                                   //image/gif
	cont6 := http.DetectContentType([]byte("MThd\x00\x00\x00\x06\x00\x01"))             //audio/midi
	fmt.Println(cont1)
	fmt.Println(cont2)
	fmt.Println(cont3)
	fmt.Println(cont4)
	fmt.Println(cont5)
	fmt.Println(cont6)
}
