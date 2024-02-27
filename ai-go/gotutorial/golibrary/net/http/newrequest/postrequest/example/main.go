/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 10:10:38
*/
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Instance struct {
	Values Values `json:"values"`
}

type Values struct {
	Company   string
	FirstName string `json:"First Name"`
	LastName  string `json:"Last Name"`
}

// 对应示例server/main.go
func main() {
	url := "http://127.0.0.1:8080/post"
	s := Instance{Values: Values{
		Company:   "123",
		FirstName: "123",
		LastName:  "123",
	}}

	js, err := json.MarshalIndent(&s, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(js))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed,err:%v\n\n", err)
		return
	}
	fmt.Println(string(b))
}
