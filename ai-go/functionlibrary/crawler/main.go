/*
@File   : main.go
@Author : pan
@Time   : 2024-08-21 10:30:22
*/
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

var downLoadPath = "本地保存路径"
var tUrl string = "网站地址"
var DownUrl string = "网站地址/wp-admin/admin-ajax.php"
var wg sync.WaitGroup

func main() {
	info := getPicInfo(tUrl)
	for k := range info {
		s := info[k]
		if len(s) != 0 {
			wg.Add(1)
			go DownloadFileProgress(strings.TrimSpace(s), strconv.Itoa(k))
		}
	}
	wg.Wait()
}

func addZero(s string) string {
	if len(s) < 2 {
		return "0" + s
	} else {
		return s
	}
}

func getPicInfo(u string) []string {
	var urls []string
	start := time.Now()
	c := colly.NewCollector(func(collector *colly.Collector) {
		extensions.RandomUserAgent(collector) // 设置随机头
		collector.Async = true
	})
	// 解析页面
	c.OnHTML("li a", func(e *colly.HTMLElement) {
		urlValue := e.Attr("href")
		if strings.Contains(urlValue, "xyz") {
			c.Visit(e.Request.AbsoluteURL(urlValue))
		}
	})

	c.OnHTML("div.main_left", func(e *colly.HTMLElement) {
		e.DOM.Each(func(i int, s *goquery.Selection) {
			text := s.Find("div.down_meta_dec").Text()
			time := s.Find("span.image-info-time").Text()
			if strings.Contains(text, "zip") {
				// 解析图片的pid
				//split := strings.Split(text, ".zip")
				//pid := split[0]s
				time = strings.TrimSpace(time)
				time = time[3:]
				s := strings.Split(time, ".")
				// 下载请求
				u := "下载地址/zip/" + s[0] + "/" + addZero(s[1]) + "/" + addZero(s[2]) + "/" + text
				//fmt.Printf("find downloadUrl: %s\n", u)
				urls = append(urls, u)
			}
		})
	})
	c.OnError(func(response *colly.Response, err error) {
		fmt.Println("错误原因：", err)
	})
	time.Sleep(1 * time.Second)
	c.Visit(u)
	c.Wait()
	fmt.Printf("花费时间:%s\n", time.Since(start))
	return urls
}

type Reader struct {
	io.Reader
	Total   int64
	Current int64
}

func (r *Reader) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)

	r.Current += int64(n)
	fmt.Printf("\r进度 %.2f%%", float64(r.Current*10000/r.Total)/100)
	return
}

func DownloadFileProgress(url, filename string) {
	r, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = r.Body.Close() }()

	f, err := os.Create(downLoadPath + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = f.Close() }()

	reader := &Reader{
		Reader: r.Body,
		Total:  r.ContentLength,
	}
	_, _ = io.Copy(f, reader)
	wg.Done()
}

// 获取加密后的下载链接，返回的值是加密的，网页中的js又经过加密所以就没有去请求下载连接了
func DownInfo(sUrl string) {
	/*
		var action string = "Post_down_ajax"
		var pid string
		var name string = "资源一下载"
		var down string = "ntXXpthwZZGmoV7Yo17JmdBfqKWhX+CbqZFnY2dlYZFpkJNnlGdmlWlnXt2coQ=="
		params := url.Values{}
		params.Set("action", action)
		params.Set("pid", pid)
		params.Set("down", down)
		params.Set("name", name)
		requestData := map[string][]string{
			"action": []string{action},
			"pid":    []string{pid},
			"down":   []string{down},
			"name":   []string{name},
		}
		c2 := colly.NewCollector()
		c2.OnRequest(func(request *colly.Request) {
			request.Headers.Set("Content-Type", "application/x-www-form-urlencoded")
		})

		err := c2.Post(sUrl, requestData)
		if err != nil {
			log.Fatal(err)
		}
	*/
}

func DemoCrawler() {
	// 假设第二页的内容是直接在 HTML 中的，通过选择器定位到第二页的内容
	res, err := http.Get("https://example.com/pages")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 假设第二页的内容是通过类选择器 .page-2 来区分的
	secondPageSelector := ".page-2"
	doc.Find(secondPageSelector).Each(func(i int, s *goquery.Selection) {
		// 在这里处理第二页的每个元素
		// 例如打印出来
		s.Each(func(_ int, sel *goquery.Selection) {
			fmt.Println(sel.Text())
		})
	})
}

func GetMovie(url string) {
	fmt.Println(url)
	//new 一个 request，再设置其header
	req, _ := http.NewRequest("GET", url, nil)
	// 设置
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1;WOW64) AppleWebKit/537.36 (KHTML,like GeCKO) Chrome/45.0.2454.85 Safari/537.36 115Broswer/6.0.3")
	req.Header.Set("Referer", "https://movie.douban.com/")
	req.Header.Set("Connection", "keep-alive")
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		panic(err)
	}
	//bodyString, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(bodyString))
	if resp.StatusCode != 200 {
		fmt.Println("err")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	//

	doc.Find("#content h1").Each(func(i int, s *goquery.Selection) {
		// name
		fmt.Println("name:" + s.ChildrenFiltered(`[property="v:itemreviewed"]`).Text())
		// year
		fmt.Println("year:" + s.ChildrenFiltered(`.year`).Text())
	})

	// #info > span:nth-child(1) > span.attrs
	director := ""
	doc.Find("#info span:nth-child(1) span.attrs").Each(func(i int, s *goquery.Selection) {
		// 导演
		director += s.Text()
		//fmt.Println(s.Text())
	})
	fmt.Println("导演:" + director)
	//fmt.Println("\n")

	pl := ""
	doc.Find("#info span:nth-child(3) span.attrs").Each(func(i int, s *goquery.Selection) {
		pl += s.Text()
	})
	fmt.Println("编剧:" + pl)

	charactor := ""
	doc.Find("#info span.actor span.attrs").Each(func(i int, s *goquery.Selection) {
		charactor += s.Text()
	})
	fmt.Println("主演:" + charactor)

	typeStr := ""
	doc.Find("#info > span:nth-child(8)").Each(func(i int, s *goquery.Selection) {
		typeStr += s.Text()
	})
	fmt.Println("类型:" + typeStr)
}

func GetToplist(url string) []string {
	var urls []string
	//new 一个 request，再设置其header
	req, _ := http.NewRequest("GET", url, nil)
	// 设置
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1;WOW64) AppleWebKit/537.36 (KHTML,like GeCKO) Chrome/45.0.2454.85 Safari/537.36 115Broswer/6.0.3")
	req.Header.Set("Referer", "https://movie.douban.com/")
	req.Header.Set("Connection", "keep-alive")
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println("===============================================================", resp.StatusCode)
	//bodyString, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(bodyString))
	if resp.StatusCode != 200 {
		fmt.Println("//////////////////////////////////////", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	doc.Find("#content div div.article ol li div div.info div.hd a").
		Each(func(i int, s *goquery.Selection) {
			// year
			fmt.Printf("%v", s)
			herf, _ := s.Attr("href")
			urls = append(urls, herf)
		})
	return urls
}

func MainRun() {
	url := "https://movie.douban.com/top250?start=0"
	urls := GetToplist(url)
	fmt.Printf("%v", urls)
	for _, url := range urls {
		GetMovie(url)
	}
}
