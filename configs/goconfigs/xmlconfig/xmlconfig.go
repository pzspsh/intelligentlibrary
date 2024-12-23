/*
@File   : xmlconfig.go
@Author : pan
@Time   : 2023-06-09 16:59:47
*/
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"golang.org/x/text/encoding/charmap"
)

type XmlConfig struct {
	Cvrfdoc xml.Name      `xml:"cvrfdoc,omitempty"`
	Vuls    Vulnerability `xml:"Vulnerability,omitempty"`
}

type Vulnerability struct {
	Title      string     `xml:"Title,omitempty"`
	Cve        string     `xml:"CVE,omitempty"`
	References References `xml:"References,omitempty"`
	Notes      Notes      `xml:"Notes,omitempty"`
}

type Notes struct {
	Note []Note `xml:"Note,omitempty"`
}

// type Note struct {
// 	Description string `xml:"Type,attr,omitempty"`
// 	Published   string `xml:"Title,attr,omitempty"`
// 	Modified    string `xml:",innerxml"`
// }

type Note struct {
	Type  string `xml:"Type,attr,omitempty"` // attr表示xml的属性
	Title string `xml:"Title,attr,omitempty"`
	Text  string `xml:",innerxml"`
}

type References struct {
	Reference []Reference `xml:"Reference,omitempty"`
}

type Reference struct {
	Url  string `xml:"URL,omitempty"`
	Desc string `xml:"Description,omitempty"`
}

// type Note struct {
// 	Description `xml:"Type,attr,omitempty"`
// 	Published   `xml:"Title,attr,omitempty"`
// 	Modified    string `xml:",innerxml"`
// }

// type Description struct {
// 	Desc string `xml:",innerxml"`
// }

func ParseXml(filepath string) (*XmlConfig, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	var config XmlConfig
	decoder := xml.NewDecoder(f)
	decoder.CharsetReader = MakeCharsetReader
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func MakeCharsetReader(charset string, input io.Reader) (io.Reader, error) {
	if charset == "ISO-8859-1" {
		return charmap.Windows1252.NewDecoder().Reader(input), nil
	}
	return nil, fmt.Errorf("unknown charset: %s", charset)
}

func WriteXml(filename string, config *XmlConfig) error {
	// data, err := xml.Marshal(config)
	data, err := xml.MarshalIndent(config, "", "\t")
	if err != nil {
		return err
	}
	data = append([]byte(xml.Header), data...)
	err = os.WriteFile(filename, data, 0666)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	config, err := ParseXml("../../config.xml")
	if err != nil {
		fmt.Printf("parse xml err:%v", err)
	}
	err = WriteXml("../../writeconfig.xml", config)
	if err != nil {
		fmt.Printf("write xml data err:%v", err)
	}
}
