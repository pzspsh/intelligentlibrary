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
	"net/url"
	"strings"
)

var URL = "http://localhost:8888/user"

func main() {
	postFormData()
	postJsonData()
	delete()
	put()
	get()
}

func postJsonData() {
	fmt.Println("------------------- post json data --------------------------")

	//数据格式化
	data := map[string]interface{}{
		"name": "alnk2",
		"age":  18,
	}
	dataStr, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	//创建一个新的post请求
	request, err := http.NewRequest("POST", URL, strings.NewReader(string(dataStr)))
	if err != nil {
		panic(err)
	}

	//请求头设置
	request.Header.Add("Authorization", "token1...")       //token
	request.Header.Add("Content-Type", "application/json") //json请求

	//发送请求到服务端
	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	//获取服务端的返回值
	b, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

}

func postFormData() {
	fmt.Println("-------------------- post form data -------------------------")

	//请求body
	//formdata请求
	urlMap := url.Values{}
	urlMap.Add("name", "alnk")
	urlMap.Add("age", "18")

	//新建请求
	request, err := http.NewRequest("POST", URL, strings.NewReader(urlMap.Encode()))
	if err != nil {
		panic(err)
	}
	fmt.Println("request.url: ", request.URL)
	fmt.Println("request.method: ", request.Method)

	//请求头部信息
	request.Header.Add("Authorization", "token1...")
	//post formdata表单请求
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	//发送请求给服务端
	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	//服务端返回数据
	b, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("服务端返回的数据: ", string(b))
}

func delete() {
	fmt.Println("-------------------- delete ------------------------------")

	//携带tonken
	//通过ID删除
	request, err := http.NewRequest("DELETE", URL, nil)
	if err != nil {
		panic(err)
	}

	//请求头部信息
	request.Header.Add("Authorization", "token1...") //token

	//url参数
	query := request.URL.Query()
	query.Add("id", "1")
	query.Add("id", "2")
	query.Add("id", "3")
	request.URL.RawQuery = query.Encode()

	//发送请求给服务端
	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	//服务端返回数据
	b, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

}

func get() {
	//正常来说，GET请求所有的参数都应该是放在url上，body请求体中是没有数据的
	//请求头可能会携带token之类的参数
	fmt.Println("--------------------get请求-------------------------")

	//新建一个GET请求
	request, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		panic(err)
	}

	//请求头部信息
	//Set时候，如果原来这一项已存在，后面的就修改已有的
	//Add时候，如果原本不存在，则添加，如果已存在，就不做任何修改
	//最终服务端获取的应该是token2
	request.Header.Set("User-Agent", "自定义浏览器1...")
	request.Header.Set("User-Agent", "自定义浏览器2...")
	//header:  map[User-Agent:[自定义浏览器2...]]

	request.Header.Add("name", "alnk")
	request.Header.Add("name", "alnk2")
	//header:  map[Name:[alnk alnk2] User-Agent:[自定义浏览器2...]]

	request.Header.Add("Authorization", "token1...") //token

	fmt.Println("header: ", request.Header)

	//url参数
	query := request.URL.Query()
	query.Add("id", "1")
	query.Add("id", "2")
	query.Add("name", "alnk")
	request.URL.RawQuery = query.Encode()
	fmt.Println("request.URL: ", request.URL)
	//request.URL:  http://localhost:8888/user?id=1&id=2&name=alnk

	//发送请求给服务端
	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	//服务端返回数据
	b, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("服务端返回的数据: ", string(b))
}

func put() {
	//通过ID修改用户数据
	fmt.Println("-------------------- put --------------------")

	//数据格式化
	data := map[string]interface{}{
		"name": "alnk2",
		"age":  18,
	}
	dataStr, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	//创建一个新的post请求
	request, err := http.NewRequest("PUT", URL, strings.NewReader(string(dataStr)))
	if err != nil {
		panic(err)
	}

	//请求头设置
	request.Header.Add("Authorization", "token1...")       //token
	request.Header.Add("Content-Type", "application/json") //json请求

	//url参数
	query := request.URL.Query()
	query.Add("id", "1")
	request.URL.RawQuery = query.Encode()

	//发送请求给服务端
	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	//服务端返回数据
	b, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("服务端返回的数据: ", string(b))
}
