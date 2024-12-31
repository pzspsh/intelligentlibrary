# Golang File Download

#### 1、git project download 
```go
/*
@File   : main.go
@Author : pan
@Time   : 2024-08-19 11:57:20
*/
package main

import (
	"bytes"
	"crypto/tls"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Options struct {
	Target         string
	DownloadUrl    string
	LocalPath      string
	IsWrit         bool
	TagsLog        string
	BranchLog      string
	AllTags        bool
	AllBranch      bool
	Master         bool
	Develop        bool
	Latest         bool
	ProxyDown      string
	Proxy          string
	SingleFileDown string // 单个文件下载
}

func ParseUrl(parseurl string) {
	u, err := url.Parse(parseurl)
	if err != nil {
		fmt.Println(err)
	}
	pathsegments := strings.Split(u.Path, "/")
	if len(pathsegments) > 0 {
		projectname := pathsegments[len(pathsegments)-1]
		fmt.Println(projectname)
	} else {
		fmt.Println("no path segments found.")
	}
}

// 查找数组的索引
func FindIndex(arr []string, value string) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}
	return -1
}

func PageParse() []string {
	filenames := []string{}
	file, err := os.Open("tags.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	doc, err := goquery.NewDocumentFromReader(file)
	doc.Find("h2").Find("a").Each(func(i int, c *goquery.Selection) {
		href := c.Text()
		if href != "" {
			fmt.Println(href)
		}
	})
	if err != nil {
		fmt.Println("branches goquery html error: ", err)
	}
	return filenames
}

func CompareVersions(newvserion, oldversion string) int {
	newvserion = strings.ReplaceAll(newvserion, "v", "")
	oldversion = strings.ReplaceAll(oldversion, "v", "")
	parts1 := strings.Split(newvserion, ".")
	parts2 := strings.Split(oldversion, ".")
	length := len(parts1)
	if len(parts2) > length {
		length = len(parts2)
	}
	for i := 0; i < length; i++ {
		var num1, num2 int
		var err error
		if i < len(parts1) {
			num1, err = strconv.Atoi(parts1[i])
			if err != nil {
				fmt.Println("the version number format is incorrect:", err)
				return 0
			}
		}
		if i < len(parts2) {
			num2, err = strconv.Atoi(parts2[i])
			if err != nil {
				fmt.Println("the version number format is incorrect:", err)
				return 0
			}
		}
		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		}
	}
	return 0
}

func SearchDir(rootdir string) ([]string, error) {
	var err error
	filenamelist := []string{}
	if err = filepath.Walk(rootdir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if !info.IsDir() {
			filename := filepath.Base(path)
			filenamelist = append(filenamelist, filename)
		}
		return nil
	}); err != nil {
		return filenamelist, err
	}
	return filenamelist, nil
}

func Download(downurl, localdir string, options *Options) error {
	var err error
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	if options.ProxyDown != "" {
		if !strings.HasSuffix(options.ProxyDown, "/") {
			options.ProxyDown = options.ProxyDown + "/"
		}
		downurl = strings.Replace(downurl, "https://codeload.", options.ProxyDown, 1)
		if strings.Contains(downurl, "zip/refs") {
			downurl = strings.Replace(downurl, "zip/refs", "archive/refs", 1)
			downurl = downurl + ".zip"
		}
	} else if options.Proxy != "" {
		proxyUrl, err := url.Parse(options.Proxy)
		if err != nil {
			return err
		}
		tr.Proxy = http.ProxyURL(proxyUrl)
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   360 * time.Second,
	}
	resp, err := client.Get(downurl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(localdir)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}

func GetGithubTags(downurl string, options *Options) (map[string]string, error) {
	var err error
	var filename string
	var nexturl string
	targeturls := map[string]string{}
	parsedUrl, err := url.Parse(downurl)
	if err != nil {
		return targeturls, err
	}
	pathsegments := strings.Split(parsedUrl.Path, "/")
	if len(pathsegments) > 0 {
		filename = pathsegments[len(pathsegments)-1]
	} else {
		return targeturls, errors.New("no path segments found")
	}
	parsedUrl.Host = "codeload.github.com"
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	if options.Proxy != "" {
		proxyUrl, err := url.Parse(options.Proxy)
		if err != nil {
			return targeturls, err
		}
		tr.Proxy = http.ProxyURL(proxyUrl)
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   360 * time.Second,
	}
	begin := 0
	for {
		resp := &http.Response{}
		if begin == 0 {
			resp, err = client.Get(downurl + "/tags")
			if err != nil {
				return targeturls, err
			}
			begin++
		} else {
			resp, err = client.Get(nexturl)
			if err != nil {
				return targeturls, err
			}
		}
		body, _ := io.ReadAll(resp.Body)
		if options.IsWrit {
			tagsfile, err := os.OpenFile(options.TagsLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return targeturls, err
			}
			_, err = tagsfile.Write(body)
			if err != nil {
				return targeturls, err
			}
		}
		if !strings.Contains(string(body), "There aren’t any releases here") {
			doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
			if err != nil {
				return targeturls, err
			}
			if options.AllTags {
				doc.Find("h2").Find("a").Each(func(i int, c *goquery.Selection) {
					versions := c.Text()
					if versions != "" {
						downloadurl := parsedUrl.String() + "/zip/refs/tags/" + versions
						targeturls[downloadurl] = filename + "-" + strings.Replace(versions, "v", "", 1) + ".zip"
					}
				})
				doc.Find("div").Find("a").Each(func(i int, c *goquery.Selection) {
					if c.Text() == "Next" {
						href, exists := c.Attr("href")
						if exists {
							nexturl = "https://github.com" + href
						}
					} else if c.Parent().Find("span").Text() == "Next" {
						nexturl = ""
					}
				})
				if nexturl == "" {
					break
				}
			} else if options.Latest {
				versions := doc.Find("h2").Find("a").Eq(0).Text()
				downloadurl := parsedUrl.String() + "/zip/refs/tags/" + versions
				targeturls[downloadurl] = filename + "-" + strings.Replace(versions, "v", "", 1) + ".zip"
				break
			}
		} else {
			err = errors.New("not tags")
			break
		}
	}
	return targeturls, err
}

func GetGithubBranches(downurl string, options *Options) (map[string]string, error) {
	var err error
	var filename string
	var nextdownurl string
	targeturls := map[string]string{}
	parsedUrl, err := url.Parse(downurl)
	if err != nil {
		return targeturls, err
	}
	pathsegments := strings.Split(parsedUrl.Path, "/")
	if len(pathsegments) > 0 {
		filename = pathsegments[len(pathsegments)-1]
	} else {
		return targeturls, errors.New("no path segments found")
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	parsedUrl.Host = "codeload.github.com"
	if options.Proxy != "" {
		proxyUrl, err := url.Parse(options.Proxy)
		if err != nil {
			return targeturls, err
		}
		tr.Proxy = http.ProxyURL(proxyUrl)
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   360 * time.Second,
	}
	begin := 0
	for {
		resp := &http.Response{}
		if begin == 0 {
			resp, err = client.Get(downurl + "/branches/all?page=1")
			if err != nil {
				return targeturls, err
			}
			begin++
		} else {
			resp, err = client.Get(nextdownurl)
			if err != nil {
				return targeturls, err
			}
		}
		if options.IsWrit {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return targeturls, err
			}
			file, err := os.OpenFile(options.BranchLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return targeturls, err
			}
			_, err = file.Write(body)
			if err != nil {
				return targeturls, err
			}
		}
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return targeturls, err
		}
		if options.AllBranch {
			doc.Find("tbody").Find("tr").Find("td").Find("div").Find("a").Each(func(i int, s *goquery.Selection) {
				branch := s.Text()
				downloadurl := parsedUrl.String() + "/zip/refs/heads/" + branch
				branch = strings.ReplaceAll(branch, "/", "-")
				branch = strings.ReplaceAll(branch, "\\", "-")
				targeturls[downloadurl] = filename + "-" + branch + ".zip"
			})
			a := doc.Find("react-app").Find("div").Find("nav").Find("div").Find("a").Last()
			if a.Text() == "Next" {
				next, exists := a.Attr("href")
				if exists {
					nextdownurl = downurl + fmt.Sprintf("/branches/all?page=%v", strings.Replace(next, "#", "", 1))
				} else {
					break
				}
			} else {
				break
			}
		} else if options.Develop || options.Master {
			var develop bool
			var master bool
			doc.Find("tbody").Find("tr").Find("td").Find("div").Find("a").Each(func(i int, s *goquery.Selection) {
				branch := s.Text()
				if options.Develop {
					if JudgmentExist([]string{"dev", "develop"}, branch) {
						downloadurl := parsedUrl.String() + "/zip/refs/heads/" + branch
						targeturls[downloadurl] = filename + "-" + branch + ".zip"
						develop = true
					}
				}
				if options.Master {
					if JudgmentExist([]string{"main", "master"}, branch) {
						downloadurl := parsedUrl.String() + "/zip/refs/heads/" + branch
						targeturls[downloadurl] = filename + "-" + branch + ".zip"
						master = true
					}
				}
			})
			a := doc.Find("react-app").Find("div").Find("nav").Find("div").Find("a").Last()
			if a.Text() == "Next" {
				next, exists := a.Attr("href")
				if exists {
					nextdownurl = downurl + fmt.Sprintf("/branches/all?page=%v", strings.Replace(next, "#", "", 1))
				} else {
					break
				}
			} else {
				break
			}
			if options.Develop && options.Master && master && develop {
				break
			} else if options.Develop && develop {
				break
			} else if options.Master && master {
				break
			}
		}
	}
	return targeturls, nil
}

func GetGithubAllFile(targeturl string, options *Options) ([]string, error) {
	var err error
	targeturls := []string{}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	if options.Proxy != "" {
		proxyUrl, err := url.Parse(options.Proxy)
		if err != nil {
			return targeturls, err
		}
		tr.Proxy = http.ProxyURL(proxyUrl)
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   360 * time.Second,
	}
	resp, err := client.Get(targeturl)
	if err != nil {
		return targeturls, err
	}
	ParseBody(resp, options)
	return targeturls, err
}

func ParseBody(resp *http.Response, options *Options) ([]string, error) {
	var err error
	targeturls := []string{}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return targeturls, err
	}
	doc.Find("table").Find("tbody").Find("tr").Find("td").Each(func(i int, s *goquery.Selection) {
		colspan, exists := s.Attr("colspan")
		if exists && colspan == "1" {
			s.Find("div").Find("div").Find("div").Find("a").Each(func(i int, v *goquery.Selection) {
				arialable, exists := v.Attr("aria-label")
				if exists {
					if strings.Contains(arialable, "File") {
						href, exists := v.Attr("href")
						href = "http://github.com" + href
						if exists {
							fmt.Println("download file url: ", href)
						}
					} else if strings.Contains(arialable, "Directory") {
						href, exists := v.Attr("href")
						if exists {
							href = "http://github.com" + href
							GetGithubAllFile(href, options)
						}
					}
				}
			})
		}
	})
	return targeturls, nil
}

func DeletedFile(filelist []string, downurls map[string]string) map[string]string {
	downurlmap := downurls
	for key, value := range downurls {
		if JudgmentExist(filelist, value) {
			delete(downurlmap, key)
		}
	}
	return downurlmap
}

func JudgmentExist(filelist []string, file string) bool {
	for _, f := range filelist {
		if f == file {
			return true
		}
	}
	return false
}

func MergeMap(dest, src map[string]string) map[string]string {
	for key, value := range src {
		dest[key] = value
	}
	return dest
}

func DownloadRun(downurls map[string]string, storagedir string, options *Options) error {
	var err error
	if _, err = os.Stat(storagedir); os.IsNotExist(err) {
		if err = os.MkdirAll(storagedir, 0755); err != nil {
			return err
		}
	} else {
		filelist, err := SearchDir(storagedir)
		if err != nil {
			return err
		}
		downurls = DeletedFile(filelist, downurls) // Deleted file
	}
	for urls, filename := range downurls {
		filenamepath := filepath.Join(storagedir, filename)
		if err = Download(urls, filenamepath, options); err != nil {
			return err
		}
	}
	return err
}

func GithubProjectRun(targets, storagedir string) error {
	var err error
	var downtarget []string
	options := &Options{}
	flag.BoolVar(&options.IsWrit, "w", false, "iswrite")
	flag.StringVar(&options.TagsLog, "tlog", "", "write file path")
	flag.StringVar(&options.BranchLog, "blog", "", "write file path")
	flag.BoolVar(&options.AllTags, "alltag", false, "download all tags")
	flag.BoolVar(&options.AllBranch, "allbranch", false, "download all branch")
	flag.BoolVar(&options.Master, "master", false, "download master branches")
	flag.BoolVar(&options.Develop, "dev", false, "download develop branches")
	flag.BoolVar(&options.Latest, "latest", false, "download latest version")
	flag.StringVar(&options.Target, "target", "", "download target url")       // 如果有多个下载目标，url之间用英文“,”隔开
	flag.StringVar(&options.DownloadUrl, "downurl", "", "download target url") // 如果有多个直接下载url，url之间用英文“,”隔开
	flag.StringVar(&options.LocalPath, "dir", "", "download file path")
	flag.StringVar(&options.Proxy, "proxy", "", "proxy download")
	flag.StringVar(&options.ProxyDown, "proxydown", "", "proxy download")
	flag.Parse()
	if options.Target != "" {
		targets = options.Target
	}
	if options.LocalPath != "" {
		storagedir = options.LocalPath
	}
	if len(targets) > 0 {
		downtarget = strings.Split(targets, ",")
		if options.IsWrit {
			if options.TagsLog == "" {
				options.TagsLog = "tags.txt"
			}
			if options.BranchLog == "" {
				options.BranchLog = "branches.txt"
			}
		}
		for _, target := range downtarget {
			var downdir string
			downurlsmap := map[string]string{}
			if target != "" {
				parsedUrl, err := url.Parse(target)
				if err != nil {
					return err
				}
				pathsegments := strings.Split(parsedUrl.Path, "/")
				downdir = filepath.Join(storagedir, pathsegments[len(pathsegments)-1])
				if options.AllTags || options.Latest {
					tagsdownloadurls, err := GetGithubTags(target, options)
					if err != nil {
						return err
					}
					if len(tagsdownloadurls) > 0 {
						downurlsmap = MergeMap(downurlsmap, tagsdownloadurls)
					}
				}
				if options.AllBranch || options.Master || options.Develop {
					branchesurls, err := GetGithubBranches(target, options)
					if err != nil {
						return err
					}
					if len(branchesurls) > 0 {
						downurlsmap = MergeMap(downurlsmap, branchesurls)
					}
				}
				if len(downurlsmap) > 0 {
					if err = DownloadRun(downurlsmap, downdir, options); err != nil {
						return err
					}
				}
			} else {
				err = errors.New("target download is empty")
			}
		}
	}
	if options.DownloadUrl != "" {
		downloadurl := strings.Split(options.DownloadUrl, ",")
		for _, downurl := range downloadurl {
			var downdir string
			downurlsmap := map[string]string{}
			if strings.Contains(downurl, "zip/refs") {
				parseurl, err := url.Parse(downurl)
				if err != nil {
					return err
				}
				pathsegments := strings.Split(parseurl.Path, "/")
				index := FindIndex(pathsegments, "zip")
				file := pathsegments[index-1]
				suffix := pathsegments[len(pathsegments)-1]
				if strings.HasPrefix(suffix, "v") {
					suffix = strings.Replace(suffix, "v", "", 1)
				}
				filename := fmt.Sprintf("%v-%v.zip", file, suffix)
				downdir = filepath.Join(storagedir, file)
				downurlsmap[downurl] = filename
				if err = DownloadRun(downurlsmap, downdir, options); err != nil {
					return err
				}
			} else if strings.Contains(downurl, "archive/refs") {
				parseurl, err := url.Parse(downurl)
				if err != nil {
					return err
				}
				pathsegments := strings.Split(parseurl.Path, "/")
				index := FindIndex(pathsegments, "archive")
				file := pathsegments[index-1]
				suffix := pathsegments[len(pathsegments)-1]
				if strings.HasPrefix(suffix, "v") {
					suffix = strings.Replace(suffix, "v", "", 1)
				}
				filename := fmt.Sprintf("%v-%v", file, suffix)
				downdir := filepath.Join(storagedir, file)
				downurlsmap[downurl] = filename
				if err = DownloadRun(downurlsmap, downdir, options); err != nil {
					return err
				}
			} else if strings.Contains(downurl, "releases/download") {
				parseurl, err := url.Parse(downurl)
				if err != nil {
					return err
				}
				pathsegments := strings.Split(parseurl.Path, "/")
				index := FindIndex(pathsegments, "releases")
				file := pathsegments[index-1]
				downdir := filepath.Join(storagedir, file)
				filename := pathsegments[len(pathsegments)-1]
				downurlsmap[downurl] = filename
				if err = DownloadRun(downurlsmap, downdir, options); err != nil {
					return err
				}
			}
		}
	}
	return err
}

func main() {
	// var err error
	downtargets := "https://github.com/projectdiscovery/nuclei/tree/dev/pkg" // 下载目标
	// catalog := "../"  // 存储的目录
	options := &Options{}
	// if err = GithubProjectRun(downtargets, catalog); err != nil {
	// 	fmt.Println("error: ", err)
	// }
	if _, err := GetGithubAllFile(downtargets, options); err != nil {
		fmt.Println("get github all file error: ", err)
	}

}

/*
example:
	https://github.com/guardicore/monkey
	https://github.com/projectdiscovery/subfinder
	https://github.com/projectdiscovery/public-bugbounty-programs
	https://github.com/PuerkitoBio/goquery
	https://github.com/polaris1119/golangweekly
	https://github.com/projectdiscovery/nuclei
Download:
	https://codeload.github.com/projectdiscovery/subfinder/zip/refs/tags/v2.6.6  // tags用该例子
	https://codeload.github.com/projectdiscovery/subfinder/zip/refs/heads/dev    // branches用该例子

query all:
	https://github.com/projectdiscovery/subfinder/branches/all
	https://github.com/projectdiscovery/subfinder/tags
	https://github.com/projectdiscovery/subfinder/tags?after=v2.5.7  // 跳转上一页
	https://github.com/projectdiscovery/subfinder/tags?after=v2.4.7  // 跳转下一页

https://github.com/projectdiscovery/subfinder/releases/download/v2.6.6/subfinder_2.6.6_windows_amd64.zip
https://github.com/projectdiscovery/subfinder/releases
https://codeload.github.com/ExpLangcn/Aopo/zip/refs/heads/master // 下载
https://github.com/projectdiscovery/public-bugbounty-programs/tags  // 没有tags，标识：There aren’t any releases here

proxy download:
	https://cors.isteed.cc/
	https://githubfast.com/
	https://gitclone.com/
	https://kkgithub.com/
	https://ghproxy.net/
	https://mirror.ghproxy.com/
	https://download.yzuu.cf/
	https://sciproxy.com/
	https://download.ixnic.net/
	https://dgithub.xyz/
代理下载连接例子如下：
	https://gh.ddlc.top/github.com/projectdiscovery/subfinder/archive/refs/tags/v2.6.6.zip
	https://gh.ddlc.top/github.com/polaris1119/golangweekly/archive/refs/heads/master.zip
	https://cors.isteed.cc/github.com/golang101/golang101/archive/refs/heads/master.zip
	https://mirror.ghproxy.com/github.com/projectdiscovery/subfinder/releases/download/v2.6.6/subfinder_2.6.6_windows_amd64.zip

之后会添加下载git的各种文件(如：pdf、md等等)：
	原始url下载：
	https://raw.githubusercontent.com/guardicore/monkey/master/codecov.yml
	https://raw.githubusercontent.com/guardicore/monkey/develop/codecov.yml
	https://raw.githubusercontent.com/projectdiscovery/nuclei/main/cmd/nuclei/main.go
	https://raw.githubusercontent.com/projectdiscovery/nuclei/dev/cmd/nuclei/main.go
	https://raw.githubusercontent.com/CVEProject/cvelistV5/main/cves/2024/0xxx/CVE-2024-0007.json
	使用代理https://mirror.ghproxy.com下载：
	https://mirror.ghproxy.com/raw.githubusercontent.com/CVEProject/cvelistV5/main/cves/2024/0xxx/CVE-2024-0007.json

获取代理的URL：
	https://github.com/XIU2/UserScript
	https://github.com/XIU2/UserScript/blob/master/GithubEnhanced-High-Speed-Download.user.js
	https://update.greasyfork.org/scripts/412245/Github%20%E5%A2%9E%E5%BC%BA%20-%20%E9%AB%98%E9%80%9F%E4%B8%8B%E8%BD%BD.user.js
*/
```

```go
var DownTargetList = []string{
	// 漏洞扫描工具
	"https://github.com/shadow1ng/fscan",
	"https://github.com/jjf012/gopoc",
	"https://github.com/Adminisme/ServerScan",
	"https://github.com/netxfly/x-crack",
	"https://github.com/hack2fun/Gscan",
	"https://github.com/k8gege/LadonGo",
	"https://github.com/aquasecurity/trivy",
	"https://github.com/JustinTimperio/gomap",
	"https://github.com/projectdiscovery/nuclei",
	"https://github.com/google/osv-scanner",
	"https://github.com/ExpLangcn/Aopo",
	"https://github.com/anchore/grype",
	"https://github.com/shadow1ng/fscan",
	"https://github.com/JKme/cube",
	"https://github.com/veo/vscan",
	"https://github.com/78778443/QingScanDesk",
	"https://github.com/netxfly/Transparent-Proxy-Scanner",
	"https://github.com/wrenchonline/glint",
	"https://github.com/selinuxG/Golin",
	"https://github.com/MiSecurity/x-patrol",
	"https://github.com/s0md3v/XSStrike",
	"https://github.com/u21h2/nacs",
	"https://github.com/inbug-team/InScan",
	"https://github.com/awake1t/HackReport",
	"https://github.com/rustgopy/RGPScan",
	"https://github.com/inbug-team/SweetBabyScan", // Red Tools 渗透测试
	"https://github.com/arminc/clair-scanner",
	"https://github.com/dwisiswant0/crlfuzz",
	"https://github.com/s0md3v/XSStrike",
	"https://github.com/quay/clair",
	"https://github.com/anchore/anchore",
	"https://github.com/sqlmapproject/sqlmap",
	"https://github.com/sullo/nikto",
	"https://github.com/aquasecurity/cloudsploit",
	"https://github.com/xmendez/wfuzz",
	"https://code.goole.com/archive/p/skipfish/source",
	"https://github.com/Arachni/arachni",
	"https://github.com/zaproxy/zaproxy",
	"https://github.com/andresriancho/w3af",
	"https://github.com/OWASP/Nettacker",
	"https://github.com/owasp-amass/amass", // 深度攻击面和资产发现
	"https://github.com/RUB-NDS/Terrapin-Scanner",
	"https://github.com/securego/gosec", // 安全检查员
	"https://github.com/zan8in/afrog",
	"https://github.com/infobyte/faraday",
	"https://github.com/lal0ne/vulnerability",
	"https://github.com/NVIDIA/garak",
	"https://github.com/cloudflare/flan",
	"https://github.com/PentestPad/subzy",
	"https://github.com/Qianlitp/crawlergo",
	"https://github.com/future-architect/vuls",
	"https://github.com/dwisiswant0/crlfuzz",
	"https://github.com/Manisso/fsociety", // 夺旗赛CTF
	"https://github.com/PentestPad/subzy",
	"https://github.com/knownsec/pocsuite3",
	"https://github.com/golang/vuln",
	"https://github.com/zema1/watchvuln",
	"https://github.com/PaytmLabs/nerve",
	"https://github.com/chaitin/SafeLine",
	"https://github.com/hahwul/dalfox",
	"https://github.com/guardicore/monkey",
	"https://github.com/awake1t/linglong",

	// 端口扫描工具
	"https://github.com/veo/vscan",
	"https://github.com/k8gege/LadonGo",
	"https://github.com/XinRoom/go-portScan",
	"https://github.com/jboursiquot/portscan",
	"https://github.com/JustinTimperio/gomap",
	"https://github.com/robertdavidgraham/masscan",
	"https://github.com/RustScan/RustScan",
	"https://github.com/GhostTroops/scan4all",
	"https://github.com/projectdiscovery/naabu",
	"https://github.com/xuxueyun-one/PortScan",
	"https://github.com/lcvvvv/gonmap",
	"https://github.com/HToTH/fuckcdn", // 全网扫描找出真实IP
	"https://github.com/s0md3v/Smap",
	"https://github.com/XinRoom/go-portScan",
	"https://github.com/AlphabugX/AScanPort",
	"https://github.com/xs25cn/scanPort",
	"https://github.com/vesche/scanless",
	"https://github.com/GhostTroops/scan4all",
	"https://github.com/se55i0n/PortScanner",
	"https://github.com/redtoolskobe/scaninfo", // 信息扫描
	"https://github.com/s0md3v/Silver",
	"https://github.com/JimYJ/scanproxy",
	"https://github.com/ATpiu/asset-scan",
	"https://github.com/4dogs-cn/TXPortMap",
	"https://github.com/lcvvvv/kscan",
	"https://github.com/selinuxG/Golin",
	"http://www.github-zh.com/collections/port-scanner",
	"https://github.com/nmap/nmap",
	"https://github.com/qq431169079/PortScanner-3",

	// 算法
	"https://github.com/jwasham/coding-interview-university",
	"https://github.com/TheAlgorithms/Python",
	"https://github.com/trekhleb/javascript-algorithms",
	"https://github.com/yangshun/tech-interview-handbook",
	"https://github.com/halfrost/LeetCode-Go",
	"https://github.com/donnemartin/interactive-coding-challenges",
	"https://github.com/keon/algorithms",
	"https://github.com/azl397985856/leetcode",
	"https://github.com/keon/algorithms",
	"https://github.com/facebook/zstd",
	"https://github.com/TheAlgorithms/C",
	"https://github.com/TheAlgorithms/Go",
	"https://github.com/twitter/the-algorithm-ml",

	// 指纹识别工具
	"https://github.com/TideSec/TideFinger",
	"https://github.com/webanalyzer/rules",
	"https://github.com/se55i0n/Webfinger",
	"https://github.com/lcvvvv/kscan",
	"https://github.com/yhy0/FuckFingerprint",
	"https://github.com/l3m0n/whatweb",
	"https://github.com/Athena1337/blackJack",
	"https://github.com/newbe3three/gotoscan",
	"https://github.com/boy-hack/goWhatweb",
	"https://github.com/lemonlove7/EHole_magic",
	"https://github.com/zhzyker/dismap",
	"https://github.com/lcvvvv/appfinger",
	"https://github.com/wintrysec/Taiji",
	"https://github.com/EdgeSecurityTeam/EHole",
	"https://github.com/EASY233/Finger",

	// 内网穿透工具
	"https://github.com/ginuerzh/gost",
	"https://github.com/cloudflare/cloudflared",
	"https://github.com/go-gost/gost",
	"https://github.com/vzex/dog-tunnel",
	"https://github.com/fatedier/frp",
	"https://github.com/ehang-io/nps",
	"https://github.com/v2fly/v2ray-core",
	"https://github.com/XTLS/Xray-core",
	"https://github.com/xtaci/kcptun",
	"https://github.com/jpillora/chisel",
	"https://github.com/p4gefau1t/trojan-go",
	"https://github.com/telepresenceio/telepresence",
	"https://github.com/traefik/traefik",
	"https://github.com/mitmproxy/mitmproxy",
	"https://github.com/SagerNet/sing-box",
	"https://github.com/nginx-proxy/nginx-proxy",
	"https://github.com/joewalnes/websocketd",
	"https://github.com/snail007/goproxy",
	"https://github.com/inconshreveable/ngrok",
	"https://github.com/antoniomika/sish",
	"https://github.com/pgrok/pgrok",
	"https://github.com/FunnyWolf/TFirewall",
	"https://github.com/Mob2003/rakshasa",
	"https://github.com/Dliv3/Venom", // 渗透测试仪的多跳代理
	"https://github.com/v2ray/v2ray-core",
	"https://github.com/v2fly/v2ray-core",
	"https://github.com/drk1wi/Modlishka",
	"https://github.com/XTLS/Xray-core",

	// 网络代理
	"https://github.com/traefik/traefik",
	"https://github.com/nginx-proxy/nginx-proxy",
	"https://github.com/oauth2-proxy/oauth2-proxy",
	"https://github.com/SagerNet/sing-box",
	"https://github.com/DNSCrypt/dnscrypt-proxy",
	"https://github.com/XX-net/XX-Net",
	"https://github.com/elazarl/goproxy",
	"https://github.com/goproxyio/goproxy",
	"https://github.com/umputun/reproxy",
	"https://github.com/malfunkt/hyperfox",
	"https://github.com/Shopify/toxiproxy",

	// 子域名扫描工具
	"https://github.com/yunxu1/dnsub",
	"https://github.com/guelfoweb/knock",
	"https://github.com/jwt1399/Sec-Tools",
	"https://github.com/knownsec/ksubdomain",
	"https://github.com/projectdiscovery/subfinder",
	"https://github.com/shadowabi/AutoDomain",
	"https://github.com/ZhuriLab/Starmap",
	"https://github.com/shadow1ng/fscan",
	"https://github.com/ExpLangcn/Aopo",
	"https://github.com/lijiejie/subDomainsBrute",
	"https://github.com/guelfoweb/knock",
	"https://github.com/yanxiu0614/subdomain3",
	"https://github.com/laramies/theHarvester",
	"https://github.com/aboul3la/Sublist3r",
	"https://github.com/evilsocket/dnssearch",
	"https://github.com/PentestPad/subzy",
	"https://github.com/haccer/subjack",
	"https://github.com/tomnomnom/assetfinder",
	"https://github.com/tomnomnom/waybackurls",
	"https://github.com/d3mondev/puredns",
	"https://github.com/PentestPad/subzy",
	"https://github.com/edoardottt/scilla",
	"https://github.com/boy-hack/ksubdomain",

	// 漏洞脚本
	"https://github.com/projectdiscovery/nuclei-templates",
	"https://github.com/trickest/cve",
	"https://github.com/nomi-sec/PoC-in-GitHub",

	// 游戏开发
	"https://github.com/viphxin/xingo",
	"https://github.com/suiyunonghen/DxTcpServer" // 游戏网络通信框架
	"https://github.com/icexin/gocraft",
	"https://github.com/hajimehoshi/ebiten",
	"https://github.com/OpenDiablo2/OpenDiablo2",

	// 运维
	"https://github.com/1Panel-dev/1Panel",

	// 大模型
	"https://github.com/ollama/ollama",
	"https://github.com/geekan/MetaGPT",
	"https://github.com/mlabonne/llm-course",
	"https://github.com/run-llama/llama_index",
	"https://github.com/QuivrHQ/quivr",
	"https://github.com/hiyouga/LLaMA-Factory",
	"https://github.com/milvus-io/milvus",
	"https://github.com/xtekky/gpt4free",
	"https://github.com/PlexPt/awesome-chatgpt-prompts-zh",
	"https://github.com/ChatGPTNextWeb/ChatGPT-Next-Web",
	"https://github.com/AUTOMATIC1111/stable-diffusion-webui",
	"https://github.com/huggingface/transformers",
	"https://github.com/pytorch/pytorch",
	"https://github.com/d2l-ai/d2l-zh",
	"https://github.com/keras-team/keras",
	"https://github.com/labmlai/annotated_deep_learning_paper_implementations",
	"https://github.com/hpcaitech/ColossalAI",
	"https://github.com/aymericdamien/TensorFlow-Examples",
	"https://github.com/fighting41love/funNLP",
	"https://github.com/Significant-Gravitas/AutoGPT",
	"https://github.com/binary-husky/gpt_academic",
	"https://github.com/geekan/MetaGPT",
	"https://github.com/2noise/ChatTTS",
	"https://github.com/Pythagora-io/gpt-pilot",
	"https://github.com/rasbt/MachineLearning-QandAI-book",
	"https://github.com/rasbt/deeplearning-models",
	"https://github.com/rasbt/machine-learning-book",
	"https://github.com/rasbt/LLM-workshop-2024",
	"https://github.com/google-research/bert",
	"https://github.com/microsoft/AI-For-Beginners",
	"https://github.com/GokuMohandas/Made-With-ML",
	"https://github.com/ymcui/Chinese-LLaMA-Alpaca",
	"https://github.com/Ciphey/Ciphey", // 夺旗赛 (CTF) 和网络安全资源#使用自然语言处理和人工智能以及一些全自动解密/解码/破解工具。
	"https://github.com/lazyprogrammer/machine_learning_examples",
	"https://github.com/amusi/Deep-Learning-Interview-Book",
	"https://github.com/meta-llama/llama",
	"https://github.com/ageitgey/face_recognition", // 人脸识别#本项目是世界上最简洁的人脸识别库，你可以使用Python和命令行工具提取、识别、操作人脸
	"https://github.com/openai/whisper",            //whisper 是一个通用语音识别模型
	"https://github.com/labring/FastGPT",
	"https://github.com/songquanpeng/one-api",
	"https://github.com/Calcium-Ion/new-api?tab=readme-ov-file",
	"https://github.com/coaidev/coai",

	// 资产扫描
	"https://github.com/maurosoria/dirsearch",
	"https://github.com/hakluke/hakrawler",
	"https://github.com/tomnomnom/assetfinder",

	// 数学
	"https://github.com/apachecn/ailearning",
	"https://github.com/3b1b/manim",
	"https://github.com/jiye-ML/math_study",
	"https://github.com/jakevdp/PythonDataScienceHandbook",
	"https://github.com/ManimCommunity/manim",

	// 机器学习
	"https://github.com/MorvanZhou/tutorials",
	"https://github.com/lawlite19/MachineLearning_Python",
	"https://github.com/shunliz/Machine-Learning",
	"https://github.com/apachecn/ailearning",

	// 深度学习
	"https://github.com/Mikoto10032/DeepLearning",
	"https://github.com/fengdu78/deeplearning_ai_books",

	// 数据库

	// 消息队列
	"https://github.com/nsqio/nsq",

	// 分布式
	"https://github.com/busgo/forest",

	// 微服务

	// 容器

	// 反编译工具

	// 爬虫工具
	"https://github.com/projectdiscovery/katana",
	"https://github.com/gocolly/colly",
	"https://github.com/scrapy/scrapy",

	// 网络安全
	"https://github.com/Clouditera/SecGPT",

	// API接口开发

	// 视频播放器
	"https://github.com/Hunlongyu/ZY-Player",
	"https://github.com/gwuhaolin/livego",

	// 爆破工具
	"https://github.com/niudaii/crack",
	"https://github.com/awake1t/PortBrute",
	"https://github.com/JKme/cube",
	"https://github.com/i11us0ry/goon",
	"https://github.com/byt3bl33d3r/CrackMapExec",
	"https://github.com/netxfly/x-crack",

	// 密码破解
	"https://github.com/Ciphey/Ciphey",

	// 人工智能

	// 机器人

	// 图形界面

	// 操作系统

	// 网络设备

	// 音视频

	// 图形图像处理

	// 自然语言处理

	// 网络通信
	"https://github.com/curl/curl",
	"https://github.com/xindong/frontd",
	"https://github.com/dxcweb/go-nat-hole",
	"https://github.com/funny/link",
	"https://github.com/gwuhaolin/lightsocks", // 网络混淆代理
	"https://github.com/txthinking/brook", // 跨平台加密功能
	"https://github.com/AdguardTeam/AdGuardHome", // 广告拦截，反跟踪dns服务器

	// 网络协议
	"https://github.com/brewlin/net-protocol",
	"https://github.com/impact-eintr/netstack",
	"https://github.com/quic-go/quic-go",
	"https://github.com/apernet/hysteria",

	// Py项目

	// Go项目

	// C项目

	// 其它
	"https://github.com/freeCodeCamp/freeCodeCamp",
	"https://github.com/EbookFoundation/free-programming-books",
	"https://github.com/sindresorhus/awesome",
	"https://github.com/public-apis/public-apis",
	"https://github.com/codecrafters-io/build-your-own-x",
	"https://github.com/jwasham/coding-interview-university",
	"https://github.com/kamranahmedse/developer-roadmap",
	"https://github.com/donnemartin/system-design-primer",
	"https://github.com/vinta/awesome-python",
	"https://github.com/practical-tutorials/project-based-learning",
	"https://github.com/TheAlgorithms/Python",
	"https://github.com/torvalds/linux",
	"https://github.com/Significant-Gravitas/AutoGPT",
	"https://github.com/microsoft/vscode",
	"https://github.com/jackfrued/Python-100-Days",
	"https://github.com/trimstray/the-book-of-secret-knowledge",
	"https://github.com/avelino/awesome-go",
	"https://github.com/ytdl-org/youtube-dl",
	"https://github.com/golang/go",
	"https://github.com/Genymobile/scrcpy",
	"https://github.com/kubernetes/kubernetes",
	"https://github.com/ollama/ollama",
	"https://github.com/syncthing/syncthing",
	"https://github.com/Lengso/iplookup", // IP反查域名
	"https://github.com/corazawaf/coraza",
	"https://github.com/kgretzky/evilginx2", // 
	"https://github.com/hashicorp/vault", // 密钥管理工具
	"https://github.com/rancher/rancher", // 管理和部署k8s完整解决方案
	"https://github.com/secdev/scapy",
	"https://github.com/hengyoush/kyanos", // 网络解析器
	"https://github.com/j3ssie/osmedeus", 
	
	
	// 夺旗赛（CTF）和网络安全资源
	"https://github.com/trimstray/the-book-of-secret-knowledge",
	"https://github.com/swisskyrepo/PayloadsAllTheThings",
	"https://github.com/sherlock-project/sherlock",
	"https://github.com/danielmiessler/SecLists",
	"https://github.com/drduh/macOS-Security-and-Privacy-Guide",
	"https://github.com/jivoi/awesome-osint",
	"https://github.com/Ciphey/Ciphey",
	"https://github.com/bettercap/bettercap",
	"https://github.com/vitalysim/Awesome-Hacking-Resources",
	"https://github.com/Gallopsled/pwntools",
	"https://github.com/gophish/gophish",
	"https://github.com/laramies/theHarvester",
	"https://github.com/s0md3v/Photon",
	"https://github.com/Manisso/fsociety",
	"https://github.com/GTFOBins/GTFOBins.github.io",
	"https://github.com/juice-shop/juice-shop",
	"https://github.com/ctf-wiki/ctf-wiki",
	"https://github.com/pwndbg/pwndbg",
	"https://github.com/google/google-ctf",
	"https://github.com/rsmusllp/king-phisher",
}
```