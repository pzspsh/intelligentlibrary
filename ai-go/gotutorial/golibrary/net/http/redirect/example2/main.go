/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 16:23:20
*/
package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/textproto"
	"net/url"
	"strconv"
	"strings"
)

func setCookieWhenRedirect(client *HttpClient, domain *url.URL, resp *Response) {

	//get cookie (JSESSIONID=MDJkZDFkMjItZTc3OS00ZWY5LTkzZWUtODI3MDBmYWFkMzg2; path=/oauth; HttpOnly)
	cookieInHeader := resp.Headers[textproto.CanonicalMIMEHeaderKey("Set-Cookie")]
	var cookieStr string
	if len(cookieInHeader) != 0 {
		cookieStr = cookieInHeader[len(cookieInHeader)-1]
	}
	if cookieStr == "" {
		return
	}
	//build cookie
	cookie := &http.Cookie{}
	parts := strings.Split(cookieStr, ";")
	for _, item := range parts {
		if item == "" {
			continue
		}
		item = strings.Trim(item, " ")
		if item == "HttpOnly" {
			cookie.HttpOnly = true
			continue
		}
		subItems := strings.Split(item, "=")
		if len(subItems) != 2 {
			continue
		}
		if subItems[0] == "Path" || subItems[0] == "path" {
			cookie.Path = subItems[1]
			continue
		}
		cookie.Name = subItems[0]
		cookie.Value = subItems[1]
	}
	// set cookie
	jar := client.CookieJar()
	if jar == nil {
		jar, _ = cookiejar.New(nil)
	}
	jar.SetCookies(domain, []*http.Cookie{cookie})
	client.SetCookieJar(jar)
	// 设置成功后下次请求domain就会带上cookie
}

type Response struct {
	Headers http.Header
	Data    []byte
}

func (resp *Response) HttpCodeIs200() (bool, error) {
	code := resp.HttpCode()
	if code != 200 {
		return false, fmt.Errorf("http-code: %d", code)
	}
	return true, nil
}

func (resp *Response) HttpCode() int {
	code := resp.Headers.Get("StatusCode")
	if code == "" {
		return 0
	}

	intCode, _ := strconv.Atoi(code)

	return intCode
}

type HttpClient struct {
	client *http.Client
}

func NewHttpClient() *HttpClient {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
	}

	return &HttpClient{client: client}
}

func (httpClient *HttpClient) CheckRedirect(fn func(req *http.Request, via []*http.Request) error) {
	httpClient.client.CheckRedirect = fn
}

func (httpClient *HttpClient) SetCookieJar(jar http.CookieJar) {
	httpClient.client.Jar = jar
}

func (httpClient *HttpClient) CookieJar() http.CookieJar {
	return httpClient.client.Jar
}

func (httpClient *HttpClient) Get(url string, header map[string]string) (*Response, error) {
	return httpClient.Call(http.MethodGet, url, header, nil)
}

func (httpClient *HttpClient) Post(url string, header map[string]string, body io.Reader) (*Response, error) {
	return httpClient.Call(http.MethodPost, url, header, body)
}

func (httpClient *HttpClient) Put(url string, header map[string]string, body io.Reader) (*Response, error) {
	return httpClient.Call(http.MethodPut, url, header, body)
}

func (httpClient *HttpClient) Delete(url string, header map[string]string) (*Response, error) {
	return httpClient.Call(http.MethodDelete, url, header, nil)
}

func (httpClient *HttpClient) Call(method, url string, header map[string]string, body io.Reader) (*Response, error) {
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	if len(header) > 0 {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	resp, err := httpClient.client.Do(req)
	if err != nil {
		return nil, err
	}

	respHeader := resp.Header

	defer resp.Body.Close()

	respCode := resp.StatusCode
	respHeader.Set("StatusCode", strconv.Itoa(respCode))
	respHeader.Set("Status", resp.Status)

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read from response'body error. response code %v. response content: %v",
			respCode, string(content))
	}

	return &Response{Headers: respHeader, Data: content}, nil
}

func main() {
	client := NewHttpClient()
	// golang http自动302跳转，但是不会在302时设置cookie
	// 所以如下代码取消自动跳转，然后手动设置cookie
	client.CheckRedirect(func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	})

	testUrl := "http://a.b.c/test"
	resp, err := client.Post(testUrl, nil, nil)

	if err != nil {
		fmt.Printf("request error '%s', cause by %v", testUrl, err)
	}

Redirect:
	if resp.HttpCode() != 200 && resp.HttpCode() != 302 {
		fmt.Printf("request error[%d] %s ", resp.HttpCode(), testUrl)
	}

	if resp.HttpCode() == 200 {
		//最后请求成功

	} else if resp.HttpCode() == 302 {
		// xWiki /oauth/login
		location := resp.Headers.Get("location")
		if location == "" {
			location = resp.Headers.Get("location")
		}
		u, err := url.Parse(testUrl)
		if err != nil {
			fmt.Printf("parse host error '%s', cause by %v", u.Host, err)
		}
		setCookieWhenRedirect(client, u, resp)
		resp, _ = client.Get(location, nil)
		//可能多次跳转
		goto Redirect
	}
}
