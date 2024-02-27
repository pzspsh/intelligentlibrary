/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 01:44:19
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//上传文件 =>multipart/form-data

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		//第一种方式
		//req.ParseMultipartForm(1024 * 1024) //解析提交内容
		//fmt.Println(req.MultipartForm)
		//file, err := req.MultipartForm.File["a"][0].Open()
		//if err != nil {
		//	fmt.Println(err)
		//}
		//io.Copy(os.Stdout, file)
		//第二种方式
		file, header, _ := req.FormFile("a")
		fmt.Println(header.Filename) //文件名
		fmt.Println(header.Size)     //文件大小
		fmt.Println(header.Header)   //文件头的相关信息
		io.Copy(os.Stdout, file)
	})
	http.ListenAndServe("127.0.0.1:8888", nil)
}
