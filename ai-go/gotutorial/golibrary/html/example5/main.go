/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 10:42:14
*/
package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

var (
	htmlTplEngine    *template.Template
	htmlTplEngineErr error
)

func init() {
	// 初始化模板引擎 并加载各层级的模板文件
	// 注意 views/* 不会对子目录递归处理 且会将子目录匹配 作为模板处理造成解析错误
	// 若存在与模板文件同级的子目录时 应指定模板文件扩展名来防止目录被作为模板文件处理
	// 然后通过 view/*/*.html 来加载 view 下的各子目录中的模板文件
	htmlTplEngine = template.New("htmlTplEngine")

	// 模板根目录下的模板文件 一些公共文件
	_, htmlTplEngineErr = htmlTplEngine.ParseGlob("views/*.html")
	if nil != htmlTplEngineErr {
		log.Panic(htmlTplEngineErr.Error())
	}
	// 其他子目录下的模板文件
	_, htmlTplEngineErr = htmlTplEngine.ParseGlob("views/*/*.html")
	if nil != htmlTplEngineErr {
		log.Panic(htmlTplEngineErr.Error())
	}

}

// index
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	_ = htmlTplEngine.ExecuteTemplate(
		w,
		"index/index",
		map[string]interface{}{"PageTitle": "首页", "Name": "sqrt_cat", "Age": 25},
	)
}

// news
func NewsHandler(w http.ResponseWriter, r *http.Request) {
	_ = htmlTplEngine.ExecuteTemplate(
		w,
		"news/index",
		map[string]interface{}{
			"PageTitle": "新闻",
			"List": []struct {
				Title     string
				CreatedAt time.Time
			}{
				{Title: "this is golang views/template example", CreatedAt: time.Now()},
				{Title: "to be honest, i don't very like this raw engine", CreatedAt: time.Now()},
			},
			"Total":  1,
			"Author": "big_cat",
		},
	)
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/index", IndexHandler)
	http.HandleFunc("/news", NewsHandler)

	serverErr := http.ListenAndServe(":8085", nil)

	if nil != serverErr {
		log.Panic(serverErr.Error())
	}
}
