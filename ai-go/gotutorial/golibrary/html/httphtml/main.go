/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 10:33:53
*/
package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	// 利用给定数据渲染模板, 并将结果写入w
	tmpl.Execute(w, "小明")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP SERVER failed,err:", err)
		return
	}
}

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func SayHe(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("./he.tmpl")
	if err != nil {
		fmt.Println("create template failed,err:", err)
		return
	}
	// 利用给定数据渲染模板, 并将结果写入w
	user := UserInfo{
		Name:   "小明",
		Gender: "男",
		Age:    18,
	}
	tmpl.Execute(w, user)
}
