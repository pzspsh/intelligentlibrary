/*
@File   : main.go
@Author : pan
@Time   : 2024-08-19 11:57:20
*/
package gitpull

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"maps"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Options struct {
	Target      string
	DownloadUrl string
	LocalPath   string
	IsWrit      bool
	TagsLog     string
	BranchLog   string
	AllTags     bool
	AllBranch   bool
	Master      bool
	Develop     bool
	Latest      bool
	ProxyDown   string
	Proxy       string
	IsDownPath  string
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

type IsDownJson struct {
	Finish     bool               `json:"finish,omitempty"`
	IsDownload map[string]DownUrl `json:"isdownload,omitempty"`
}

type DownUrl struct {
	Finish   bool   `json:"finish,omitempty"`
	Filename string `json:"filename,omitempty"`
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
			fmt.Println(downurl, "not tags")
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
	return slices.Contains(filelist, file)
}

func MergeMap(dest, src map[string]string) map[string]string {
	maps.Copy(dest, src)
	return dest
}

func GetIsDownload(file string) (map[string]IsDownJson, error) {
	var err error
	var isDownload = make(map[string]IsDownJson)
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return isDownload, err
	}
	if err = json.Unmarshal(data, &isDownload); err != nil { // 解析JSON数据
		fmt.Println("Error parsing JSON: ", err)
		return isDownload, err
	}
	return isDownload, nil
}

func WriteIsDownload(file string, isDownload map[string]IsDownJson) error {
	var err error
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	data, err := json.Marshal(isDownload)
	if err != nil {
		return err
	}
	f.Write(data)
	return nil
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

func GithubProjectRun(targets, storagedir string, options *Options, isdownload map[string]IsDownJson) error {
	var err error
	var isexec bool
	var isfull = false
	var downtarget []string
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
		for !isfull {
			var flag = true
			for _, target := range downtarget {
				downmap := make(map[string]DownUrl)
				downurlsmap := make(map[string]string)
				if target != "" {
					if _, ok := isdownload[target]; !ok {
						fmt.Printf("start get target: [%s] download url....\n", target)
						if options.AllTags || options.Latest {
							tagsdownloadurls, err := GetGithubTags(target, options)
							if err != nil {
								flag = false
								fmt.Println("get tags download url error: ", err)
								continue
							}
							if len(tagsdownloadurls) > 0 {
								downurlsmap = MergeMap(downurlsmap, tagsdownloadurls)
							}
						}
						if options.AllBranch || options.Master || options.Develop {
							branchesurls, err := GetGithubBranches(target, options)
							if err != nil {
								flag = false
								fmt.Println("get github branches error: ", err)
								continue
							}
							if len(branchesurls) > 0 {
								downurlsmap = MergeMap(downurlsmap, branchesurls)
							}
						}
						if len(downurlsmap) > 0 {
							for URL, filename := range downurlsmap {
								downmap[URL] = DownUrl{Finish: false, Filename: filename}
							}
							isdownload[target] = IsDownJson{Finish: false, IsDownload: downmap}
							isexec = true
						}
					}
				} else {
					err = errors.New("target download is empty")
				}
			}
			if flag {
				isfull = true
			}
		}
		for {
			var isexecfinish = true
			if len(isdownload) > 0 && isexec {
				for target, v1 := range isdownload {
					var finish = true
					var downdir string
					if !v1.Finish {
						parsedUrl, err := url.Parse(target)
						if err != nil {
							return err
						}
						pathsegments := strings.Split(parsedUrl.Path, "/")
						downdir = filepath.Join(storagedir, pathsegments[len(pathsegments)-1])
						for URL, v2 := range v1.IsDownload {
							var downurlmap = make(map[string]string)
							if !v2.Finish {
								fmt.Printf("target URL: [%s] start Downloading...\n", URL)
								downurlmap[URL] = v2.Filename
								if len(downurlmap) > 0 {
									if err = DownloadRun(downurlmap, downdir, options); err != nil {
										finish = false
										continue
									} else {
										v2.Finish = true
										v1.IsDownload[URL] = v2
									}
								}
							}
						}
					}
					if finish {
						v1.Finish = true
						isdownload[target] = v1
					} else {
						isexecfinish = false
						isdownload[target] = v1
					}
				}
				if isexecfinish {
					break
				} else {
					time.Sleep(30 * time.Second)
				}
			} else {
				break
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
