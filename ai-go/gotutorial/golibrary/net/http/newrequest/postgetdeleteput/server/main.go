/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 12:22:24
*/
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func main() {
	http.HandleFunc("/user", userHandleFunc)
	http.ListenAndServe(":8888", nil)

}

func userHandleFunc(w http.ResponseWriter, r *http.Request) {
	//post 请求
	if r.Method == "POST" {
		//form data
		if r.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
			fmt.Println("r.body: ", r.Body)
			fmt.Println("name: ", r.FormValue("name"))
			fmt.Println("age: ", r.FormValue("age"))
			w.Write([]byte("post form data is ok!\n"))
			return
		}

		//json
		if r.Header.Get("Content-Type") == "application/json" {
			b, err := io.ReadAll(r.Body)
			if err != nil {
				w.Write([]byte("post json data is error!\n"))
				return
			}
			fmt.Println("post json: ", string(b))
			w.Write([]byte("post json data is ok!\n"))
		}

	}

	//delete 请求
	if r.Method == "DELETE" {
		queryParams := r.URL.Query()
		requestStr, err := json.Marshal(queryParams)
		if err != nil {
			w.Write([]byte("delete is error\n"))
			return
		}
		fmt.Println("delete: ", string(requestStr))
		w.Write([]byte("delete is ok!\n"))
		return
	}

	//put 请求
	if r.Method == "PUT" {
		//id
		queryParams := r.URL.Query()
		requestStr, err := json.Marshal(queryParams)
		if err != nil {
			w.Write([]byte("put is error\n"))
			return
		}
		fmt.Println("put id: ", string(requestStr))

		b, err := io.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("put json data is error!\n"))
			return
		}
		fmt.Println("put json: ", string(b))

		w.Write([]byte("put is ok!\n"))
		return
	}

	//get 请求
	if r.Method == "GET" {
		//正常来说，GET请求所有的参数都应该是放在url上，body请求体中是没有数据的
		//请求头可能会携带token之类的参数
		url := r.URL
		fmt.Println("请求的url: ", url)

		//获取url中的所有参数
		queryParma := r.URL.Query()
		jsonStr, err := json.Marshal(queryParma)
		if err != nil {
			w.Write([]byte("get is error!\n"))
			return
		}
		fmt.Println("获取所有的参数，jsonStr: ", string(jsonStr))
		fmt.Println("获取单个参数： ", r.URL.Query()["name"])

		fmt.Println("获取所有的请求头head: ", r.Header)
		fmt.Println("获取指定的请求头head token: ", r.Header.Get("Authorization"))

		w.Write([]byte("get is ok!\n"))
		return
	}

}
