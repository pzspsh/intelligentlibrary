/*
@File   : main.go
@Author : pan
@Time   : 2023-06-06 15:06:32
*/
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type DataInfo struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Tags string `json:"tags,omitempty"`
}

func main() {
	datapath := "./ai-go/functionlibrary/parsejson/data.json"
	f, err := os.Open(datapath)
	if err != nil {
		fmt.Println("os Open error:", err)
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("io ReadAll error:", err)
	}
	var a map[string]DataInfo
	if err = json.Unmarshal(data, &a); err != nil {
		fmt.Println("json Unmarshal error:", err)
	}
	fmt.Println(a)
}
