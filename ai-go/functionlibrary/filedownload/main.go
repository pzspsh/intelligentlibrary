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
	"path/filepath"
	"strings"

	gitpull "function/filedownload/githubpull"
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

func DownRun(downloadurl, loadpath string) error {
	var err error
	u, err := url.Parse(downloadurl)
	if err != nil {
		return err
	}
	filename := path.Base(u.Path) // 获取URL请求文件名
	loadpathfile := loadpath + filename
	ext := filepath.Ext(filename)
	file := strings.TrimSuffix(filename, ext)
	fmt.Println(filename)     // dev.zip
	fmt.Println(file)         // dev
	fmt.Println(loadpathfile) // path/filepath/dev.zip
	err = DownloadFile(loadpathfile, downloadurl)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func GitDownload() error {
	var err error
	downtargets := "" // 下载目标
	catalog := "../"  // 存储的目录
	if err = gitpull.GithubProjectRun(downtargets, catalog); err != nil {
		fmt.Println("github download error: ", err)
	}
	return err
}

func main() {
	/*
		loadpath := "path/filepath/"
		downloadUrl := `https://github.com/projectdiscovery/dnsx/archive/refs/heads/dev.zip`
		if err := DownRun(downloadUrl, loadpath); err != nil {
			fmt.Println("download error: ", err)
		}
	*/
	if err := GitDownload(); err != nil {
		fmt.Println("github download error: ", err)
	}
}
