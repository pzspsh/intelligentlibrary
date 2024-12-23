/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 16:20:54
*/
package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	go server()
	time.Sleep(time.Second)
	mUrl := "http://127.0.0.1:12345/post"
	{ // 常规方法
		req, err := http.NewRequest(http.MethodPost, mUrl, nil)
		if err != nil {
			log.Fatal(err)
		}
		resp, err := http.DefaultClient.Do(req)
		if resp != nil {
			defer resp.Body.Close()
		}
		if err != nil {
			log.Fatal(err)
		}
		byt, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp.StatusCode, "|", string(byt[:128]))
	}

	{ // 去掉自动处理重定向
		req, err := http.NewRequest(http.MethodPost, mUrl, nil)
		if err != nil {
			log.Fatal(err)
		}
		resp, err := http.DefaultTransport.RoundTrip(req)
		if resp != nil {
			defer resp.Body.Close()
		}
		if err != nil {
			log.Fatal(err)
		}
		byt, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp.StatusCode, "|", string(byt[:128]))
	}

	{ // 另一种不要重定向的方法
		req, err := http.NewRequest(http.MethodPost, mUrl, nil)
		if err != nil {
			log.Fatal(err)
		}
		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse /* 不进入重定向 */
			},
		}
		resp, err := client.Do(req)
		if resp != nil {
			defer resp.Body.Close()
		}
		if err != nil {
			log.Fatal(err)
		}
		byt, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp.StatusCode, "|", string(byt[:128]))
	}
}

// 下面开启一个服务，重定向到百度
func server() {
	http.HandleFunc("/post", mPost)
	http.ListenAndServe(":12345", nil)
}

func mPost(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://www.baidu.com", http.StatusFound)
	w.Write([]byte(time.Now().String()))
}
