/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:32:08
*/
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	//说明写入PostForm中的字段name-value,field2有两个value，一个作为Form，一个作为PostForm
	//filename指明写入的文件
	//注意value1、value2和binary data上面的空行不能够删除，否则会报错：malformed MIME header initial line
	postData :=
		`--xxx
Content-Disposition: form-data; name="field1" 

value1
--xxx
Content-Disposition: form-data; name="field2"

value2
--xxx
Content-Disposition: form-data; name="file"; filename="file"
Content-Type: application/octet-stream
Content-Transfer-Encoding: binary

binary data
--xxx--
`
	req := &http.Request{
		Method: "POST",
		Header: http.Header{"Content-Type": {`multipart/form-data; boundary=xxx`}},
		Body:   io.NopCloser(strings.NewReader(postData)), //NopCloser用一个无操作的Close方法包装输入参数然后返回一个ReadCloser接口
	}

	initialFormItems := map[string]string{
		"language": "Go",
		"name":     "gopher",
		"skill":    "go-ing",
		"field2":   "initial-value2",
	}

	req.Form = make(url.Values) //url.Values即map[string][]string
	for k, v := range initialFormItems {
		req.Form.Add(k, v)
	}

	//应该解析的位置是这里，否则会导致最终的结果与构想的结果不同
	err := req.ParseMultipartForm(10000)
	if err != nil {
		fmt.Printf("unexpected multipart error %v\n", err)
	}

	fmt.Println(req.Form)     //map[language:[Go] name:[gopher] skill:[go-ing] field2:[initial-value2 value2] field1:[value1]]
	fmt.Println(req.PostForm) //map[field1:[value1] field2:[value2]]
	//本字段只有在调用ParseMultipartForm后才有效
	fmt.Println(req.MultipartForm) //&{map[field1:[value1] field2:[value2]] map[file:[0xc00009e140]]}

	// file, fileHeader, err := req.FormFile("field1") // 出错，返回：
	// 2019/02/13 18:40:02 http: no such file
	// exit status 1
	file, fileHeader, err := req.FormFile("file") // 成功
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(file)       //{0xc0000a3080}
	fmt.Println(fileHeader) //&{file map[Content-Transfer-Encoding:[binary] Content-Disposition:[form-data; name="file"; filename="file"] Content-Type:[application/octet-stream]] 11 [98 105 110 97 114 121 32 100 97 116 97] }
}
