/*
@File   : main.go
@Author : pan
@Time   : 2024-11-01 13:16:19
*/
package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type DataJson struct {
	Python     DataInfo `json:"python,omitempty"`
	C          DataInfo `json:"c,omitempty"`
	Go         DataInfo `json:"go,omitempty"`
	Cpp        DataInfo `json:"cpp,omitempty"`
	Javascript DataInfo `json:"javascript,omitempty"`
	Typescript DataInfo `json:"typescript,omitempty"`
	Vue        DataInfo `json:"type,omitempty"`
}

type DataInfo struct {
	Page  int
	Stars int
	Url   []string
}

type Options struct {
	Proxy    string `json:"proxy,omitempty"`
	Language string `json:"language,omitempty"`
	Stars    int    `json:"stars,omitempty"`
	Page     int    `json:"page,omitempty"`
}

func (o *Options) GetGithubUrl() string {
	/*
		https://github.com/search?q=language:python+stars:>=100&type=repositories&p=1
		https://github.com/search?q=language:go+stars:>=100&type=repositories&p=2
		https://github.com/search?q=language:python+stars:>=100&type=repositories&s=stars&o=desc&p=1
	*/
	// return fmt.Sprintf("https://github.com/search?q=language:%s+stars:>=%d&type=repositories&p=%d", o.Language, o.Stars, o.Page)
	return `https://github.com/search?q=language%3Ago+stars%3A%3E%3D100&type=repositories&p=2`
}

func (o *Options) GitCrawler() error {
	var err error
	var GithubUrl string
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	if o.Proxy != "" {
		proxyurl, err := url.Parse(o.Proxy)
		if err != nil {
			return err
		}
		tr.Proxy = http.ProxyURL(proxyurl)
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   360 * time.Second,
	}
	GithubUrl = o.GetGithubUrl()
	resp, err := client.Get(GithubUrl)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return err

}

func Parse() error {
	var err error

	return err
}

func main() {
	var err error
	opt := &Options{
		Language: "python",
		Stars:    100,
		Page:     1,
	}
	if err = opt.GitCrawler(); err != nil {
		fmt.Println()
	}
}
