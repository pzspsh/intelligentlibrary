/*
@File   : main.go
@Author : pan
@Time   : 2024-10-29 10:46:02
*/
package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

// Note 结构体表示XML数据
type Note struct {
	XMLName xml.Name `xml:"Note"`
	Type    string   `xml:"Type,attr"`
	Ordinal int      `xml:"Ordinal,attr"`
	Content string   `xml:",chardata"`
}

func main() {
	// XML数据
	data := `<Note Type="Description" Ordinal="1">ip_input.c in BSD-derived TCP/IP implementations allows remote attackers to cause a denial of service (crash or hang) via crafted packets.</Note>`

	// 创建一个空的Note结构体
	var note Note

	// 解析XML数据
	err := xml.Unmarshal([]byte(data), &note)
	if err != nil {
		log.Fatalf("Error while unmarshalling XML: %v", err)
	}

	// 检查Type属性是否为"Description"
	if note.Type == "Description" {
		fmt.Println("Content:", note.Content)
	} else {
		fmt.Println("The Type attribute is not 'Description'")
	}
}