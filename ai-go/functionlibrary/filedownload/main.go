/*
@File   : main.go
@Author : pan
@Time   : 2023-06-06 14:59:03
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

func GitDownload() error {
	var err error
	var downurlstr string
	var downurllist []string
	downname := []string{"nuclei-templates", "awesome-search-queries", "docs", "proxify", "nuclei", "tlsx", "aix", "httpx", "gcache", "naabu", "notify", "ratelimit", "utils", "mapcidr", "goflags", "tinydns", "useragent", "shuffledns", "asnmap", "cdncheck", "pdtm", "subfinder", "retryabledns", "retryablehttp-go", "katana", "chaos-client", "tldfinder", "cvemap", "rawhttp", "alterx", "public-bugbounty-programs", "gologger", "cloudlist", "hmap", "interactsh-web", "clistats", "dnsx", "interactsh", "dsl", "fastdialer", "wappalyzergo", "uncover", "nuclei-action", "freeport", "networkpolicy", "actions", "goleak", "ldapserver", "ipranger", "openrisk", "templates-stats", "fuzzing-templates", "sarif", "gozero", "machineid", "martian", "gostruct", "go-smb2", "simplehttpserver", "nuclei-ai-extension", "wallpapers", "httpx-action", "tailwindcss", "js-proto-docs", "yamldoc-go", "goconfig", "blackrock", "sslcert", "roundrobin", "nuclei-docs", "eslint-config", "fdmax", "sqlc-go-builder", "nvd", "asyncsqs", "n3iwf", "mapsutil", "stringsutil", "js-yaml-source-map", "filekv", "network-fingerprint", "rdap", "cloudlist-action", "fasttemplate", "smb", "iputil", "fileutil", "reflectutil", "httputil", "cryptoutil", "folderutil", "urlutil", "executil", "sliceutil", "sqlc-builder", "notify-action", "naabu-action", "subfinder-action", "dnsx-action", "collaborator", "pd-actions", "dnsprobe", "resolvercache-go", "expirablelru", "urlfinder", "tunnelx", "nuclei-templates-ai"}
	projectdiscovery := "https://github.com/projectdiscovery/" // 下载目标
	for _, name := range downname {
		downurllist = append(downurllist, projectdiscovery+name)
	}
	githubURL := []string{"https://github.com/Threekiii/Awesome-POC", "https://github.com/sqlmapproject/sqlmap", "https://github.com/containrrr/watchtower", "https://github.com/future-architect/vuls", "https://github.com/swisskyrepo/PayloadsAllTheThings", "https://github.com/aquasecurity/trivy", "https://github.com/The-Art-of-Hacking/h4cker", "https://github.com/chaitin/SafeLine", "https://github.com/anchore/grype", "https://github.com/google/osv-scanner", "https://github.com/shadow1ng/fscan", "https://github.com/fatedier/frp", "https://github.com/traefik/traefik", "https://github.com/mitmproxy/mitmproxy", "https://github.com/ehang-io/nps", "https://github.com/v2fly/v2ray-core", "https://github.com/XTLS/Xray-core", "https://github.com/SagerNet/sing-box", "https://github.com/snail007/goproxy", "https://github.com/Shopify/toxiproxy", "https://github.com/lqqyt2423/go-mitmproxy", "https://github.com/guardicore/monkey", "https://github.com/MrWQ/vulnerability-paper", "https://github.com/Qianlitp/crawlergo"}
	if len(githubURL) > 0 {
		downurllist = append(downurllist, githubURL...)
	}
	downurlstr = strings.Join(downurllist, ",")
	catalog := "../" // 存储的目录
	if err = gitpull.GithubProjectRun(downurlstr, catalog); err != nil {
		fmt.Println("github download error: ", err)
	}
	return err
}

func GitHubProjectsDownload(filepath string) error {
	var err error
	var catalog = "../"
	var downurllist string
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
	if err = gitpull.GithubProjectRun(downurllist, catalog); err != nil {
		fmt.Println("github download error: ", err)
	}
	return err
}

func main() {
	// loadpath := ""
	// // downloadUrl := `https://cve.mitre.org/data/downloads/allitems-cvrf.xml`
	// downloadUrl := `https://raw.githubusercontent.com/CVEProject/cvelistV5/main/cves/2024/0xxx/CVE-2024-0007.json`
	// if err := DownRun(downloadUrl, loadpath); err != nil {
	// 	fmt.Println("download error: ", err)
	// }

	if err := GitDownload(); err != nil { // go run main.go -dir /path/folder -master -dev -latest
		fmt.Println("github download error: ", err)
	}

	// filepath := ""
	// GitHubProjectsDownload(filepath)
}
