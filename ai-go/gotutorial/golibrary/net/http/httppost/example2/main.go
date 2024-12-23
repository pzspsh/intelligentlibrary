/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 10:31:26
*/
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func httpPost(requestUrl string) (err error) {
	data := make(map[string]interface{})
	data["name"] = "seemmo"
	data["passwd"] = "da123qwe"
	jsonData, _ := json.Marshal(data)

	resp, err := http.Post(requestUrl, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		fmt.Printf("get request failed, err:[%s]", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyContent, err := io.ReadAll(resp.Body)
	fmt.Printf("resp status code:[%d]\n", resp.StatusCode)
	fmt.Printf("resp body data:[%s]\n", string(bodyContent))
	return
}

// 服务实例server2/main.go
func main() {
	var url = "http://10.10.19.200:8000/index"
	httpPost(url)
}
