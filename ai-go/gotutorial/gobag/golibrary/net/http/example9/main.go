/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 01:44:19
*/
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//上传文件 =>multipart/form-data

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		//io.Copy(os.Stdout, req.Body) //读取请求头的数据,流只能解析一次
		decoder := json.NewDecoder(req.Body) //解析流
		var info map[string]interface{}
		decoder.Decode(&info) //将解析的coder对象放在map中（里面可以放任意接口）
		fmt.Println(info)
		fmt.Println(info["a"])
	})
	http.ListenAndServe("127.0.0.1:8888", nil)
}
