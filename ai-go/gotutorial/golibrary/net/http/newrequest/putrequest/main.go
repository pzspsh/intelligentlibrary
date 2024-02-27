/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 11:09:26
*/
package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	targetUrl := "https://b959e645-00ae-4bc3-8a55-7224d08b1d91.mock.pstmn.io/user/1"
	payload := strings.NewReader("{\"name\":\"张瑀楠\"}")
	req, _ := http.NewRequest("PUT", targetUrl, payload)
	req.Header.Add("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println(response)
}
