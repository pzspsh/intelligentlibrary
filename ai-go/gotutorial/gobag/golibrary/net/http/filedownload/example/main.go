/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 10:19:13
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {
	http.HandleFunc("/", downloadHandler) //   设置访问路由
	http.ListenAndServe(":8080", nil)
}
func downloadHandler(w http.ResponseWriter, r *http.Request) {
	fileName := "ceshi.png"  //filename  文件名
	path := "./data/images/" //文件存放目录防止用户下载其他目录文件
	file, err := os.Open(path + fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	fileNames := url.QueryEscape(fileName) // 防止中文乱码
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("Content-Disposition", "attachment; filename=\""+fileNames+"\"")
	if err != nil {
		fmt.Println("Read File Err:", err.Error())
	} else {
		w.Write(content)
	}
}
