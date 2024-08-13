/*
@File   : main.go
@Author : pan
@Time   : 2023-06-06 14:59:03
*/
package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
)

func DownloadFile(filepath string, url string) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}

func main() {
	loadpath := "path/filepath/"
	downloadUrl := `https://github.com/projectdiscovery/dnsx/archive/refs/heads/dev.zip`
	u, err := url.Parse(downloadUrl)
	if err != nil {
		panic(err)
	}
	filename := path.Base(u.Path) // 获取URL请求文件名
	loadpathfile := loadpath + filename
	fmt.Println(loadpathfile)
	err = DownloadFile(loadpathfile, downloadUrl)
	if err != nil {
		fmt.Println(err)
	}
}
