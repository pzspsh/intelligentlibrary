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
	"os"
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
	downloadUrl := `http://ip:port/path/test.zip`
	err := DownloadFile(loadpath, downloadUrl)
	if err != nil {
		fmt.Println(err)
	}
}
