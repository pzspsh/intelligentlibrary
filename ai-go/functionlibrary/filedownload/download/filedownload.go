/*
@File   : filedownload.go
@Author : pan
@Time   : 2023-06-06 14:37:21
*/
package filedownload

import (
	"crypto/tls"
	"io"
	"net/http"
	"os"
)

/*
文件下载：
1、网络请求文件下载
2、。。。
*/

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
