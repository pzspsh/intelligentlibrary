/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:47:45
*/
package main

import (
	"fmt"
	"io"
	"net/http"
)

func Test() {
	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.SetBasicAuth("test", "123456")
	fmt.Println(req.Proto)
	cookie := &http.Cookie{
		Name:  "test",
		Value: "12",
	}
	req.AddCookie(cookie)                     //添加cookie
	fmt.Println(req.Cookie("test"))           //获取cookie
	fmt.Println(req.Cookies())                //获取cookies
	req.Header.Set("User-Agent", "useragent") //设定ua
	fmt.Println(req.UserAgent())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		content, _ := io.ReadAll(resp.Body)
		fmt.Println(string(content))
	}

}
func main() {
	Test()
}
