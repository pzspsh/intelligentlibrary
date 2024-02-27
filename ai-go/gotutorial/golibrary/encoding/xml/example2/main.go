/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 16:50:20
*/
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

type Recurlyservers struct { //后面的内容是struct tag，标签，是用来辅助反射的
	XMLName     xml.Name `xml:"servers"`      //将元素名写入该字段
	Version     string   `xml:"version,attr"` //将version该属性的值写入该字段
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"` //Unmarshal函数直接将对应原始XML文本写入该字段
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func main() {
	file, err := os.Open("servers.xml")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)

	fmt.Printf("XMLName: %#v\n", v.XMLName)
	fmt.Printf("Version: %q\n", v.Version)

	fmt.Printf("Server: %v\n", v.Svs)
	for i, svs := range v.Svs {
		fmt.Println(i)
		fmt.Printf("Server XMLName: %#v\n", svs.XMLName)
		fmt.Printf("Server ServerName: %q\n", svs.ServerName)
		fmt.Printf("Server ServerIP: %q\n", svs.ServerIP)
	}
	fmt.Printf("Description: %q\n", v.Description)
}
