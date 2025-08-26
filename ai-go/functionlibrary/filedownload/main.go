/*
@File   : main.go
@Author : pan
@Time   : 2023-06-06 14:59:03
*/
package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
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

type DownloadData struct {
	Page  int               `json:"page,omitempty"`
	Stars int               `json:"stars,omitempty"`
	Urls  map[string]string `json:"url,omitempty"`
}

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

// func GitDownload() error {
// 	var err error
// 	downtargets := "https://github.com/infobyte/faraday" // 下载目标
// 	catalog := "../"                                     // 存储的目录
// 	if err = gitpull.GithubProjectRun(downtargets, catalog); err != nil {
// 		fmt.Println("github download error: ", err)
// 	}
// 	return err
// }

var options = &gitpull.Options{}

func GitHubProjectsDownload(filepath string) error {
	var err error
	var catalog = "../"
	var downurllist string
	var isdownload = make(map[string]gitpull.IsDownJson)
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()
	content, err := io.ReadAll(file) // 读取文件内容
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}
	var datas map[string]DownloadData
	if err = json.Unmarshal(content, &datas); err != nil { // 解析JSON数据
		fmt.Println("Error parsing JSON:", err)
		return err
	}

	for _, vs := range datas {
		for url := range vs.Urls {
			if downurllist == "" {
				downurllist = url
			} else {
				downurllist = downurllist + "," + url
			}
		}
	}
	if options.IsDownPath != "" {
		if isdownload, err = gitpull.GetIsDownload(options.IsDownPath); err != nil {
			return err
		}
	}
	if err = gitpull.GithubProjectRun(downurllist, catalog, options, isdownload); err != nil {
		fmt.Println("github download error: ", err)
	}
	return err
}

func GitDownload() error {
	var err error
	var downurlstr string
	var isdownload = make(map[string]gitpull.IsDownJson)
	var downurllist []string
	/* abcdefghijklmnopqrstuvwxyz */
	downname := []string{
		"aix", "asnmap", "alterx", "awesome-search-queries", "actions", "asyncsqs",
		"blackrock",
		"cdncheck", "cvemap", "cryptoutil", "collaborator", "clistats", "cloudlist", "chaos-client", "cloudlist-action", "cleanhttp",
		"docs", "dsl", "dnsprobe", "dnsx-action", "dnsx",
		"eslint-config", "expirablelru", "executil",
		"fileutil", "freeport", "filekv", "folderutil", "fdmax", "fasttemplate", "fastdialer", "fuzzing-templates",
		"gcache", "gologger", "goflags", "gozero", "gostruct", "go-smb2", "goconfig", "goleak",
		"httpx", "httpx-action", "httputil", "hmap",
		"interactsh-web", "interactsh", "ipranger", "iputil",
		"js-yaml-source-map", "js-proto-docs",
		"katana",
		"ldapserver",
		"mapsutil", "machineid", "mapcidr", "martian",
		"nuclei", "nuclei-templates-ai", "nuclei-templates", "naabu", "notify", "nuclei-action", "nuclei-ai-extension", "networkpolicy", "n3iwf", "nvd", "network-fingerprint", "nuclei-docs", "notify-action", "naabu-action",
		"openrisk",
		"proxify", "pdtm", "public-bugbounty-programs", "pd-actions",
		"retryablehttp-go", "ratelimit", "retryabledns", "rawhttp", "roundrobin", "rdap", "resolvercache-go", "reflectutil",
		"shuffledns", "subfinder", "simplehttpserver", "sarif", "sqlc-go-builder", "stringsutil", "sliceutil", "smb", "sslcert", "sqlc-builder", "subfinder-action",
		"tlsx", "tinydns", "tldfinder", "tailwindcss", "templates-stats", "tunnelx",
		"utils", "useragent", "uncover", "urlutil", "urlfinder",
		"wappalyzergo", "wallpapers",
		"yamldoc-go",
	}
	projectdiscovery := "https://github.com/projectdiscovery/" // 下载目标
	for _, name := range downname {
		downurllist = append(downurllist, projectdiscovery+name)
	}
	githubURL := []string{"https://github.com/Threekiii/Awesome-POC", "https://github.com/sqlmapproject/sqlmap", "https://github.com/containrrr/watchtower", "https://github.com/future-architect/vuls", "https://github.com/swisskyrepo/PayloadsAllTheThings", "https://github.com/aquasecurity/trivy", "https://github.com/The-Art-of-Hacking/h4cker", "https://github.com/chaitin/SafeLine", "https://github.com/anchore/grype", "https://github.com/google/osv-scanner", "https://github.com/shadow1ng/fscan", "https://github.com/fatedier/frp", "https://github.com/traefik/traefik", "https://github.com/mitmproxy/mitmproxy", "https://github.com/ehang-io/nps", "https://github.com/v2fly/v2ray-core", "https://github.com/XTLS/Xray-core", "https://github.com/SagerNet/sing-box", "https://github.com/snail007/goproxy", "https://github.com/Shopify/toxiproxy", "https://github.com/lqqyt2423/go-mitmproxy", "https://github.com/guardicore/monkey", "https://github.com/MrWQ/vulnerability-paper", "https://github.com/Qianlitp/crawlergo"}
	githubURL = []string{}
	if len(githubURL) > 0 {
		downurllist = append(downurllist, githubURL...)
	}
	downurlstr = strings.Join(downurllist, ",")
	catalog := "../" // 存储的目录
	if options.IsDownPath != "" {
		if isdownload, err = gitpull.GetIsDownload(options.IsDownPath); err != nil {
			return err
		}
	}
	if err = gitpull.GithubProjectRun(downurlstr, catalog, options, isdownload); err != nil {
		return err
	} else if options.IsDownPath != "" {
		if err = gitpull.WriteIsDownload(options.IsDownPath, isdownload); err != nil {
			return err
		}
	}
	return err
}

func getParams() {
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
	flag.StringVar(&options.IsDownPath, "isdown", "", "is download path")
	flag.Parse()
}

func GitDownloadRun() {
	var err error
	getParams()
	if options.LocalPath != "" {
		fmt.Println("github download start ...  1")
		if err = GitDownload(); err != nil { // go run main.go -dir /path/folder -master -dev -latest -alltag
			fmt.Println("github download error: ", err)
		}
	} else {
		fmt.Println("github download start ...  2")
		options.Master = true
		options.Develop = true
		options.Latest = true
		// options.AllTags = true
		options.IsDownPath = "githubpull/isdownload.json"
		options.LocalPath = "/path/folder" // /home/datas/2025备份/20250623备份
		if err = GitDownload(); err != nil {
			fmt.Println("github download success")
		}
	}
}

func main() {
	// loadpath := ""
	// // downloadUrl := `https://cve.mitre.org/data/downloads/allitems-cvrf.xml`
	// downloadUrl := `https://raw.githubusercontent.com/CVEProject/cvelistV5/main/cves/2024/0xxx/CVE-2024-0007.json`
	// if err := DownRun(downloadUrl, loadpath); err != nil {
	// 	fmt.Println("download error: ", err)
	// }
	// filepath := ""
	// GitHubProjectsDownload(filepath)
	GitDownloadRun()
}
