/*
@File   : main.go
@Author : pan
@Time   : 2024-08-05 14:43:10
*/
package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Url := "https://metabase.livesorted.com"
	// Url := "https://mb2.yrdhxm.org.cn:1443"
	// Url := "http://static-immersion-book.iqiyi.com"
	// Url := "http://akm506.inter.iqiyi.com"
	// Url := "http://*c5cnepgsm37t41ni7pl0.sso.ndrc.gov.cn"
	// Url := "https://你好.pas2323##nzhonsdfsdre23g.com.&^%$:443/hello.html"
	// Url := "mb2.yrdhxm.org.cn:1443"
	// Url := "ftp://mb2.yrdhxm.org.cn:23"
	// Url := "https://10.0.35.74:10000/?hello.html"
	// Url := "https://pan:zhosdf$ng@example.com"
	// Url := "http://mb2.yrdhxm.org.cn潘:1443" // err Url
	Url := "http://example.com/page#section1"
	r := regexp.MustCompile("((http|https|.*?)://|)(((\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}|[a-zA-Z0-9\u4e00-\u9fa5-@!#$&*%?^+=~#.]+)(:(\\d+)|)))")
	if r.MatchString(Url) {
		res := r.FindAllStringSubmatch(Url, -1)[0]
		aid := res[3]
		asset := res[5]
		port := res[7]
		protocol := res[2]
		fmt.Println(res)
		fmt.Println(aid)
		fmt.Println(asset)
		fmt.Println(port)
		fmt.Println(protocol)
	} else {
		fmt.Println("url错误")
	}
}
