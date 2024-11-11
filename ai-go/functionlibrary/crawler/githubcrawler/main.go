/*
@File   : main.go
@Author : pan
@Time   : 2024-11-01 13:16:19
*/
package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"

	// "strconv"
	"time"
)

var (
	regex  = regexp.MustCompile(`<a class="prc-Link-Link-85e08" href="(.*?)".*?<a class="Box-sc-g0xbh4-0 iPuHRc prc-Link-Link-85e08" href=".*?" aria-label="(.*?)"`)
	pagere = regexp.MustCompile(`<a href=".*?p=\d+" aria-label="Page \d+" class="Pagination__Page-sc-cp45c9-0 gnHNlv">(\d+)</a>`)
)

var DataJson = make(map[string]DataInfo)

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
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0")
	req.Header.Set("cookie", `_octo=GH1.1.1287721027.1723187765; preferred_color_mode=dark; tz=Asia%2FShanghai; _device_id=f57b67404efe5cc6a70dab7f97adc7f0; saved_user_sessions=56292066%3AVywVjvgmOYXIlmdCnIBjEzN1t8gjUrcYW9ajKesrMjdd_M8M; tz=Asia%2FShanghai; color_mode=%7B%22color_mode%22%3A%22auto%22%2C%22light_theme%22%3A%7B%22name%22%3A%22light%22%2C%22color_mode%22%3A%22light%22%7D%2C%22dark_theme%22%3A%7B%22name%22%3A%22dark%22%2C%22color_mode%22%3A%22dark%22%7D%7D; user_session=VywVjvgmOYXIlmdCnIBjEzN1t8gjUrcYW9ajKesrMjdd_M8M; __Host-user_session_same_site=VywVjvgmOYXIlmdCnIBjEzN1t8gjUrcYW9ajKesrMjdd_M8M; logged_in=yes; dotcom_user=pzspsh; _gh_sess=bMlXYCXpVULgQ5qbJXoVJaEbGvwFrNgmJXxZqKmHzDElhRK2ZJ%2BmfAMcMlKg1jTvQ%2Fe9aZLKACSR91x5ehtyQFYPBfiaM7wH5cVDK3AxWb8MESVQNLGwKnrRPkWfY93sY0v0fZ7FNMJrCmEAZoVLjLpstpCD%2FzwWuJy1BohpkfSNyGRO%2FOSlPEqsxxdEVH1INEvXqasNVQOnedlFIneVSS8IIqz3eX4huBEwAfLeU5SXYczkMEpUeCE4ylJ1j69%2FBFyT4J%2FjRSxseTAmQ3mv75LktGLMaJO7oOMQsAtyenPSd5oRa%2B1pQ%2Fx56illwXrUQx2T31VHZQ%2FyrYKB6i%2FLtNGid%2B5Mf%2BFBOi%2FFdj5m8HqLupn9QYlPB3PvR83wf7DC550KRUQGM96oZ5prvim3VFLeuTSP8cypipHym0XyaqnMWR0s--6nD2wfj%2FFM%2Foait7--sTtPfuFgVi57uI%2FcIH5eJQ%3D%3D`)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
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
			data.Page = o.Page
			data.Stars = o.Stars
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
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0")
	req.Header.Set("cookie", `_octo=GH1.1.1287721027.1723187765; preferred_color_mode=dark; tz=Asia%2FShanghai; _device_id=f57b67404efe5cc6a70dab7f97adc7f0; saved_user_sessions=56292066%3AVywVjvgmOYXIlmdCnIBjEzN1t8gjUrcYW9ajKesrMjdd_M8M; tz=Asia%2FShanghai; color_mode=%7B%22color_mode%22%3A%22auto%22%2C%22light_theme%22%3A%7B%22name%22%3A%22light%22%2C%22color_mode%22%3A%22light%22%7D%2C%22dark_theme%22%3A%7B%22name%22%3A%22dark%22%2C%22color_mode%22%3A%22dark%22%7D%7D; user_session=VywVjvgmOYXIlmdCnIBjEzN1t8gjUrcYW9ajKesrMjdd_M8M; __Host-user_session_same_site=VywVjvgmOYXIlmdCnIBjEzN1t8gjUrcYW9ajKesrMjdd_M8M; logged_in=yes; dotcom_user=pzspsh; _gh_sess=bMlXYCXpVULgQ5qbJXoVJaEbGvwFrNgmJXxZqKmHzDElhRK2ZJ%2BmfAMcMlKg1jTvQ%2Fe9aZLKACSR91x5ehtyQFYPBfiaM7wH5cVDK3AxWb8MESVQNLGwKnrRPkWfY93sY0v0fZ7FNMJrCmEAZoVLjLpstpCD%2FzwWuJy1BohpkfSNyGRO%2FOSlPEqsxxdEVH1INEvXqasNVQOnedlFIneVSS8IIqz3eX4huBEwAfLeU5SXYczkMEpUeCE4ylJ1j69%2FBFyT4J%2FjRSxseTAmQ3mv75LktGLMaJO7oOMQsAtyenPSd5oRa%2B1pQ%2Fx56illwXrUQx2T31VHZQ%2FyrYKB6i%2FLtNGid%2B5Mf%2BFBOi%2FFdj5m8HqLupn9QYlPB3PvR83wf7DC550KRUQGM96oZ5prvim3VFLeuTSP8cypipHym0XyaqnMWR0s--6nD2wfj%2FFM%2Foait7--sTtPfuFgVi57uI%2FcIH5eJQ%3D%3D`)
	resp, err := client.Do(req)
	if err != nil {
		return body, err
	}
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}
	return body, err
}

func Parse(body []byte, result map[string]string) (map[string]string, error) {
	var err error
	datalist := regex.FindAllSubmatch(body, -1)
	if len(datalist) > 0 {
		for _, value := range datalist {
			url := "https://github.com/" + string(value[1])
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
		for key, value := range map2 {
			map1[key] = value
		}
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
	opt := &Options{
		Language: "python",
		Stars:    100,
		Page:     1,
		Jsonfile: "data.json",
	}
	if err = opt.GitCrawler(); err != nil {
		fmt.Println()
	}
}
