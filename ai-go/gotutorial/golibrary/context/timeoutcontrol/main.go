/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 16:59:48
*/
package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "https://www.baidu.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	// 构造一个超时间为50毫秒的Context
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	req = req.WithContext(ctx)

	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	out, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(out))
}
