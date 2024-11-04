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
	return fmt.Sprintf("https://github.com/search?q=language:%s+stars:>=%d&type=repositories&p=%d", o.Language, o.Stars, o.Page)
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
	fmt.Println("BBBBBBB", string(body))
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
