/*
@File   : main.go
@Author : pan
@Time   : 2024-01-19 17:11:19
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

/*
一、什么是 CDN？
CDN 全称 Content Delivery Network ，中文名为内容分发网络。它是建立在现有互联网之上的一个智能虚拟网络，
利用网络中的边缘节点，把网站的内容和数据分发到用户最近的节点，从而加速用户访问网站的速度。

二、CDN 的工作原理
	CDN 的工作原理非常简单，其主要步骤如下：
		用户向 CDN 服务器发送请求；
		CDN 服务器收到请求后，先判断请求的内容是否在当前节点上，如果已缓存则直接返回相应结果；
		如果请求的内容未缓存，则 CDN 服务器向源服务器（一般指网站服务器）发起请求；
		源服务器将相应内容传输给 CDN 服务器，并缓存一份到 CDN 服务器；
		CDN 服务器将内容返回给客户端。
	由此可见，CDN 的优势主要在于就近访问、负载均衡和缓存技术。

三、使用 golang 实现 CDN 的基本流程
	在使用 golang 实现 CDN 的过程中，我们需要完成以下几个基本步骤：
		接收客户端请求，解析请求，确定请求的具体文件路径；
		判断文件是否存在于 CDN 节点缓存中，如存在则直接返回给客户端；
		如果不存在缓存，则向原始服务器请求文件，同时复制文件到缓存目录，之后返回给客户端；
		不断地处理客户端和服务器的请求，实现高并发、高性能的响应能力。

四、golang 实现 CDN 的代码
	在使用 golang 实现 CDN 代码之前，我们需要安装相应的依赖，例如 Gorm 等库。接下来我们将展示
golang 实现 CDN 的代码具体实现，代码如下所示：
*/

func main() {
	// 静态文件服务器地址和缓存目录
	StaticServer := "http://server.static.com/"
	cacheDir := "./cache/"

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// 解析客户端请求路径
		path := strings.TrimPrefix(request.URL.Path, "/")
		path = strings.Replace(path, "/", "-", -1)

		// 检查是否存在缓存文件
		cacheFile := cacheDir + path
		data, err := os.ReadFile(cacheFile)
		if err == nil {
			writer.Write(data)
			return
		}

		// 如果缓存文件不存在，则向静态服务器请求文件
		fileUrl := StaticServer + request.URL.Path
		response, err := http.Get(fileUrl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "http.Get() error: %v", err)
			writer.WriteHeader(404)
			return
		}

		// 将静态文件复制到缓存目录，并返回给客户端
		defer response.Body.Close()
		data, err = io.ReadAll(response.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ReadAll() error: %v", err)
			writer.WriteHeader(500)
			return
		}
		os.WriteFile(cacheFile, data, 0644)
		writer.Write(data)
	})
	http.ListenAndServe(":8080", nil)
}
