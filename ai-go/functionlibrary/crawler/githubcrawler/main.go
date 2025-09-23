package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"maps"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"time"
)

var (
	DataJson = make(map[string]DataInfo)
	cookie   string
	regex    = regexp.MustCompile(`<a class="prc-Link-Link-85e08" href="(.*?)".*?<a class="Box-sc-g0xbh4-0 iPuHRc prc-Link-Link-85e08" href=".*?" aria-label="(.*?)"`)
	pagere   = regexp.MustCompile(`<a href=".*?p=\d+" aria-label="Page \d+" class=".*?">(\d+)</a>`)
)

type DataInfo struct {
	Page  int               `json:"page,omitempty"`
	Stars int               `json:"stars,omitempty"`
	Url   map[string]string `json:"url,omitempty"`
}

type Options struct {
	Proxy    string `json:"proxy,omitempty"`
	Language string `json:"language,omitempty"`
	Stars    int    `json:"stars,omitempty"`
	Page     int    `json:"page,omitempty"`
	Jsonfile string
}

func (o *Options) GetGithubUrl() string {
	/*
		https://github.com/search?q=algorithm+stars%3A%3E%3D2000&type=repositories&s=stars&o=desc&p=1 //  搜索algorithm stars大于等于2000
		https://github.com/search?q=algorithm&type=repositories&s=stars&o=desc&p=2 // 搜索algorithm
		https://github.com/search?q=algorithm&type=repositories&s=stars&o=desc // 搜索algorithm(算法)
		https://github.com/search?q=leetcode&type=repositories&s=stars&o=desc  // 搜索leetcode
		https://github.com/search?q=language:python+stars:>=100&type=repositories&p=1
		https://github.com/search?q=language:go+stars:>=100&type=repositories&p=2
		https://github.com/search?q=language:python+stars:>=100&type=repositories&s=stars&o=desc&p=1
		https://github.com/search?q=language%3Apython+stars%3A100..2000&stars%3C1000=&type=repositories&s=stars&o=desc&p=100 // 搜索stars数100到2000直接
	*/
	return fmt.Sprintf("https://github.com/search?q=language:%s+stars:>=%d&type=repositories&p=%d", o.Language, o.Stars, o.Page)
}

func (o *Options) GitCrawler() error {
	var err error
	var page int
	var GithubUrl string
	var datajson = DataJson
	var result = make(map[string]string)
	if datajson, err = ParseJson(o.Jsonfile); err != nil {
		return err
	}
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
	req, _ := http.NewRequest("GET", GithubUrl, nil)
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36 Edg/135.0.0.0")
	req.Header.Set("Cookie", cookie)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	resp.Body.Close()
	pagedata := pagere.FindAllSubmatch(body, -1)
	if len(pagedata) > 0 {
		nlen := len(pagedata)
		pagestr := pagedata[nlen-1][1]
		page, err = strconv.Atoi(string(pagestr))
		if err != nil {
			return err
		}
	}
	if result, err = Parse(body, result); err != nil {
		return err
	}
	fmt.Println("page: ", page)
	if page > 0 {
		for i := o.Page; i <= page; i++ {
			o.Page = i + 1
			body, err = o.GetGithubBody()
			if err != nil {
				fmt.Println("get github body error: ", err)
				continue
			}
			if result, err = Parse(body, result); err != nil {
				fmt.Println("parse github body error: ", err)
				continue
			}
		}
	}

	if len(result) > 0 {
		if _, ok := datajson[o.Language]; ok {
			data := datajson[o.Language]
			data.Url = MergeMap(data.Url, result)
			datajson[o.Language] = data
		} else {
			datajson[o.Language] = DataInfo{
				Page:  o.Page,
				Stars: o.Stars,
				Url:   result,
			}
		}
	}
	if err = WriteJson(o.Jsonfile, datajson); err != nil {
		return err
	}
	return err

}

func (o *Options) GetGithubBody() ([]byte, error) {
	var err error
	var body []byte
	var GithubUrl string
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	if o.Proxy != "" {
		proxyurl, err := url.Parse(o.Proxy)
		if err != nil {
			return body, err
		}
		tr.Proxy = http.ProxyURL(proxyurl)
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   360 * time.Second,
	}
	GithubUrl = o.GetGithubUrl()
	req, _ := http.NewRequest("GET", GithubUrl, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36 Edg/135.0.0.0")
	req.Header.Set("Cookie", cookie)
	resp, err := client.Do(req)
	if err != nil {
		return body, err
	}
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()
	return body, err
}

func Parse(body []byte, result map[string]string) (map[string]string, error) {
	var err error
	datalist := regex.FindAllSubmatch(body, -1)
	if len(datalist) > 0 {
		for _, value := range datalist {
			url := "https://github.com" + string(value[1])
			stars := string(value[2])
			result[url] = stars
		}
	}
	return result, err
}

func ParseJson(file string) (map[string]DataInfo, error) {
	var err error
	var datajson = DataJson
	if _, err = os.Stat(file); os.IsNotExist(err) {
		if _, err = os.Create(file); err != nil {
			return datajson, err
		} else {
			return datajson, nil
		}
	}
	data, err := os.ReadFile(file)
	if err != nil {
		return datajson, err
	}
	if len(data) == 0 {
		return datajson, err
	}
	if err = json.Unmarshal(data, &datajson); err != nil {
		return datajson, err
	}
	return datajson, err
}

func MergeMap(map1, map2 map[string]string) map[string]string {
	if len(map1) > 0 {
		maps.Copy(map1, map2)
		return map1
	} else {
		return map2
	}
}

func WriteJson(file string, datainfo map[string]DataInfo) error {
	var err error
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	data, err := json.Marshal(&datainfo)
	if err != nil {
		return err
	}
	f.Write(data)
	return err
}

func main() {
	var err error
	cookie = ``
	opt := &Options{Language: "python", Stars: 100, Page: 1, Jsonfile: "data.json"}
	if err = opt.GitCrawler(); err != nil {
		fmt.Println()
	}
}
