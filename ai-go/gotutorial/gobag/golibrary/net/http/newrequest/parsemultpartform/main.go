/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:27:17
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

// func main() {
// 	//说明写入PostForm中的字段name-value,field2有两个value，一个作为Form，一个作为PostForm
// 	//注意value1、value2和binary data上面的空行不能够删除，否则会报错：malformed MIME header initial line
// 	postData :=
// 		`--xxx
// Content-Disposition: form-data; name="field1"

// value1
// --xxx
// Content-Disposition: form-data; name="field2"

// value2
// --xxx
// Content-Disposition: form-data; name="file"; filename="file"
// Content-Type: application/octet-stream
// Content-Transfer-Encoding: binary

// binary data
// --xxx--
// `
// 	req := &http.Request{
// 		Method: "POST",
// 		Header: http.Header{"Content-Type": {`multipart/form-data; boundary=xxx`}},
// 		Body:   io.NopCloser(strings.NewReader(postData)), //NopCloser用一个无操作的Close方法包装输入参数然后返回一个ReadCloser接口
// 	}
// 	// req.ParseForm()//在POST表单中，这种解析是没有用的，要使用下面的，当然，这个只是为了查看目前的表单值，其实不应该在这里解析
// 	err := req.ParseMultipartForm(10000)
// 	if err != nil {
// 		fmt.Printf("unexpected multipart error %v\n", err)
// 	}
// 	fmt.Println(req)
// 	fmt.Println(req.Body)
// 	fmt.Println(req.Form)     //map[field1:[value1] field2:[value2]]
// 	fmt.Println(req.PostForm) //map[field1:[value1] field2:[value2]]，现在两者值是相同的，但是下面req.Form.Add后就变了
// 	fmt.Println()

// 	initialFormItems := map[string]string{
// 		"language": "Go",
// 		"name":     "gopher",
// 		"skill":    "go-ing",
// 		"field2":   "initial-value2",
// 	}

// 	req.Form = make(url.Values) //url.Values即map[string][]string
// 	for k, v := range initialFormItems {
// 		req.Form.Add(k, v)
// 	}

// 	//应该解析的位置是这里，否则会导致最终的结果与构想的结果不同
// 	// err := req.ParseMultipartForm(10000)
// 	// if err != nil {
// 	//     fmt.Printf("unexpected multipart error %v\n", err)
// 	// }
// 	fmt.Println(req)
// 	fmt.Println(req.Body)
// 	fmt.Println(req.Form)     //map[language:[Go] name:[gopher] skill:[go-ing] field2:[initial-value2]]
// 	fmt.Println(req.PostForm) //map[field1:[value1] field2:[value2]]

// 	wantForm := url.Values{
// 		"language": []string{"Go"},
// 		"name":     []string{"gopher"},
// 		"skill":    []string{"go-ing"},
// 		"field1":   []string{"value1"},
// 		"field2":   []string{"initial-value2", "value2"},
// 	}
// 	if !reflect.DeepEqual(req.Form, wantForm) { //这里会报出不相等的结果
// 		fmt.Printf("req.Form = %v, want %v\n", req.Form, wantForm)
// 	}

// 	wantPostForm := url.Values{
// 		"field1": []string{"value1"},
// 		"field2": []string{"value2"},
// 	}
// 	if !reflect.DeepEqual(req.PostForm, wantPostForm) {
// 		fmt.Printf("req.PostForm = %v, want %v\n", req.PostForm, wantPostForm)
// 	}
// }

// 正确为：
func main() {
	//说明写入PostForm中的字段name-value,field2有两个value，一个作为Form，一个作为PostForm
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
	fmt.Println(req)
	fmt.Println(req.Body)
	fmt.Println(req.Form)     //map[language:[Go] name:[gopher] skill:[go-ing] field2:[initial-value2 value2] field1:[value1]]
	fmt.Println(req.PostForm) //map[field1:[value1] field2:[value2]]

	wantForm := url.Values{
		"language": []string{"Go"},
		"name":     []string{"gopher"},
		"skill":    []string{"go-ing"},
		"field1":   []string{"value1"},
		"field2":   []string{"initial-value2", "value2"},
	}
	if !reflect.DeepEqual(req.Form, wantForm) {
		fmt.Printf("req.Form = %v, want %v\n", req.Form, wantForm)
	}

	wantPostForm := url.Values{
		"field1": []string{"value1"},
		"field2": []string{"value2"},
	}
	if !reflect.DeepEqual(req.PostForm, wantPostForm) {
		fmt.Printf("req.PostForm = %v, want %v\n", req.PostForm, wantPostForm)
	}
}
