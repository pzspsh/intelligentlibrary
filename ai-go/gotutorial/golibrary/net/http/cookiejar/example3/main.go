/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 16:14:01
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

func main() {
	jar, _ := cookiejar.New(nil)
	var cookies []*http.Cookie
	cookie := &http.Cookie{
		Name:   "M_WEIBOCN_PARAMS",
		Value:  "rl%3D1",
		Path:   "/",
		Domain: ".weibo.cn",
	}
	cookies = append(cookies, cookie)
	cookie = &http.Cookie{
		Name:   "SUB",
		Value:  "xxx",
		Path:   "/",
		Domain: ".weibo.cn",
	}
	cookies = append(cookies, cookie)
	cookie = &http.Cookie{
		Name:   "_T_WM",
		Value:  "xxx",
		Path:   "/",
		Domain: ".weibo.cn",
	}
	cookies = append(cookies, cookie)
	cookie = &http.Cookie{
		Name:   "gsid_CTandWM",
		Value:  "xxx",
		Path:   "/",
		Domain: ".weibo.cn",
	}
	cookies = append(cookies, cookie)
	u, _ := url.Parse("http://weibo.cn/search/?vt=4")
	jar.SetCookies(u, cookies)
	fmt.Println(jar.Cookies(u))
	client := &http.Client{
		Jar: jar,
	}
	postData := url.Values{}
	postData.Set("keyword", "尹相杰")
	postData.Set("smblog", "搜微博")
	req, _ := http.NewRequest("POST", "http://weibo.cn/search/?vt=4", strings.NewReader(postData.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		panic(nil)
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(string(body))
}
