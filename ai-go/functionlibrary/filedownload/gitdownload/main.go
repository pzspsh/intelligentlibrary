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
	DownUrl   string
	DownPath  string
	IsWrit    bool
	TagsLog   string
	BranchLog string
	AllTags   bool
	AllBranch bool
	Master    bool
	Develop   bool
	Latest    bool
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

func CompareVersions(v1, v2 string) int {
	v1 = strings.ReplaceAll(v1, "v", "")
	v2 = strings.ReplaceAll(v2, "v", "")
	parts1 := strings.Split(v1, ".")
	parts2 := strings.Split(v2, ".")
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
				fmt.Println("版本号格式错误:", err)
				return 0
			}
		}
		if i < len(parts2) {
			num2, err = strconv.Atoi(parts2[i])
			if err != nil {
				fmt.Println("版本号格式错误:", err)
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

func Download(downurl, localdir string) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
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
	var filename string
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
	client := &http.Client{
		Transport: tr,
		Timeout:   360 * time.Second,
	}
	respbranches, err := client.Get(downurl + "/branches/all")
	if err != nil {
		return targeturls, err
	}
	if options.IsWrit {
		branchesbody, err := io.ReadAll(respbranches.Body)
		if err != nil {
			return targeturls, err
		}
		branchesfile, err := os.OpenFile(options.BranchLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return targeturls, err
		}
		_, err = branchesfile.Write(branchesbody)
		if err != nil {
			return targeturls, err
		}
	}
	doc, err := goquery.NewDocumentFromReader(respbranches.Body)
	if err != nil {
		return targeturls, err
	}
	doc.Find("tbody").Find("tr").Find("td").Find("div").Find("a").Each(func(i int, s *goquery.Selection) {
		branch := s.Text()
		if branch != "" {
			if options.AllBranch {
				downloadurl := parsedUrl.String() + "/zip/refs/heads/" + branch
				branch = strings.ReplaceAll(branch, "/", "-")
				branch = strings.ReplaceAll(branch, "\\", "-")
				targeturls[downloadurl] = filename + "-" + branch + ".zip"
			} else if options.Develop && options.Master {
				downbranch := []string{"dev", "develop", "main", "master"}
				if JudgmentExist(downbranch, branch) {
					downloadurl := parsedUrl.String() + "/zip/refs/heads/" + branch
					targeturls[downloadurl] = filename + "-" + branch + ".zip"
				}
			} else if options.Master {
				master := []string{"main", "master"}
				if JudgmentExist(master, branch) {
					downloadurl := parsedUrl.String() + "/zip/refs/heads/" + branch
					targeturls[downloadurl] = filename + "-" + branch + ".zip"
				}
			} else if options.Develop {
				develop := []string{"dev", "develop"}
				if JudgmentExist(develop, branch) {
					downloadurl := parsedUrl.String() + "/zip/refs/heads/" + branch
					targeturls[downloadurl] = filename + "-" + branch + ".zip"
				}
			}
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

func DownloadRun(downurls map[string]string, storagedir string) error {
	var err error
	for urls, filename := range downurls {
		filenamepath := filepath.Join(storagedir, filename)
		if err = Download(urls, filenamepath); err != nil {
			return err
		}
	}
	return err
}

func GithubProjectRun(targets, storagedir string) error {
	var err error
	options := &Options{}
	flag.BoolVar(&options.IsWrit, "w", false, "iswrite")
	flag.StringVar(&options.TagsLog, "tlog", "", "write file path")
	flag.StringVar(&options.BranchLog, "blog", "", "write file path")
	flag.BoolVar(&options.AllTags, "alltag", false, "download all tags")
	flag.BoolVar(&options.AllBranch, "allbranch", false, "download all branch")
	flag.BoolVar(&options.Master, "master", false, "download master branches")
	flag.BoolVar(&options.Develop, "dev", false, "download develop branches")
	flag.BoolVar(&options.Latest, "latest", false, "download latest version")
	flag.StringVar(&options.DownUrl, "target", "", "download target url")
	flag.StringVar(&options.DownPath, "dir", "", "download file path")
	flag.Parse()
	var downtarget []string
	if options.DownUrl != "" {
		targets = options.DownUrl
	}
	downtarget = strings.Split(targets, ",")
	if options.IsWrit {
		if options.TagsLog == "" {
			options.TagsLog = "tags.txt"
		}
		if options.BranchLog == "" {
			options.BranchLog = "branches.txt"
		}
	}
	if options.DownPath != "" {
		storagedir = options.DownPath
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
				if _, err = os.Stat(downdir); os.IsNotExist(err) {
					if err = os.MkdirAll(downdir, 0755); err != nil {
						return err
					}
				} else {
					filelist, err := SearchDir(downdir)
					if err != nil {
						return err
					}
					downurlsmap = DeletedFile(filelist, downurlsmap) // 删选文件
				}
				if err = DownloadRun(downurlsmap, downdir); err != nil {
					return err
				}
			}
		} else {
			err = errors.New("target download url error")
		}
	}
	return err
}

func main() {
	var err error
	downtargets := "https://github.com/projectdiscovery/nuclei,https://github.com/projectdiscovery/subfinder" // 下载目标
	catalog := "../"                                                                                          // 存储的目录
	if err = GithubProjectRun(downtargets, catalog); err != nil {
		fmt.Println("error: ", err)
	}
}

/*
例如：
	https://github.com/projectdiscovery/subfinder
	https://github.com/projectdiscovery/public-bugbounty-programs
	https://github.com/PuerkitoBio/goquery
下载：
	https://codeload.github.com/projectdiscovery/subfinder/zip/refs/tags/v2.6.6  // tags用该例子
	https://codeload.github.com/projectdiscovery/subfinder/zip/refs/heads/dev    // branches用该例子

查询所有：
	https://github.com/projectdiscovery/subfinder/branches/all
	https://github.com/projectdiscovery/subfinder/tags
	https://github.com/projectdiscovery/subfinder/tags?after=v2.5.7  // 跳转上一页
	https://github.com/projectdiscovery/subfinder/tags?after=v2.4.7  // 跳转下一页

https://github.com/projectdiscovery/subfinder/releases/download/v2.6.6/subfinder_2.6.6_windows_amd64.zip
https://github.com/projectdiscovery/subfinder/releases
https://codeload.github.com/ExpLangcn/Aopo/zip/refs/heads/master // 下载
https://github.com/projectdiscovery/public-bugbounty-programs/tags  // 没有tags，标识：There aren’t any releases here

还未支持添加代理下：
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
	https://cors.isteed.cc/github.com/golang101/golang101/archive/refs/heads/master.zip
*/
