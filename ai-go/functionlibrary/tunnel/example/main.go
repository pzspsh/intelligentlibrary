/*
@File   : main.go
@Author : pan
@Time   : 2024-12-17 11:24:03
*/
package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"sync"
)

var (
	addr     = "127.0.0.1:7100"
	username = "administrator"
	password = "1234567"
)

// tunnel 通道处理
func tunnel(w http.ResponseWriter, r *http.Request) {
	//判断请求方法
	if r.Method != http.MethodConnect {
		log.Println(r.Method, r.RequestURI)
		http.NotFound(w, r) //404
		return
	}
	//获取用户名与密码
	auth := r.Header.Get("Proxy-Authorization") //获取客户端授权信息
	//设置用户名与密码
	r.Header.Set("Authorization", auth)

	//验证账户密码
	u, p, ok := r.BasicAuth() //BasicAuth依赖Authorization
	if !ok || !(username == u || password == p) {
		log.Printf("bad credential: username %s or password %s\n", u, p)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	//获取目标服务器地址
	dstAddr := r.RequestURI

	//连接远程服务器
	dstConn, err := net.Dial("tcp", dstAddr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer dstConn.Close()

	//为客户端返回成功消息
	w.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))

	//劫持writer获取潜在conn
	//HTTP是应用层协议，下层TCP是网络层协议，hijack可从HTTP Response获取TCP连接，若是HTTPS服务器则是TLS连接。
	//bio是带缓冲的读写者
	srcConn, bio, err := w.(http.Hijacker).Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer srcConn.Close()
	//创建两个线程
	wg := &sync.WaitGroup{}
	wg.Add(2)
	//并发执行单元1: 将TCP连接拷贝到HTTP连接中
	go func() {
		defer wg.Done()
		//缓存处理
		n := bio.Reader.Buffered()
		if n > 0 {
			n64, err := io.CopyN(dstConn, bio, int64(n))
			if n64 != int64(n) || err != nil {
				log.Printf("io.CopyN: %d %v\n", n64, err)
				return
			}
		}
		//进行全双工的双向数据拷贝（中继）
		io.Copy(dstConn, srcConn) //relay: src->dst
	}()
	//并发执行单元2：将HTTP连接拷贝到TCP连接中
	go func() {
		defer wg.Done()
		//进行全双工的双向数据拷贝（中继）
		io.Copy(srcConn, dstConn) //relay:dst->src
	}()
	wg.Wait()
}

// 服务器 go run main.go
// 客户端 curl -p --proxy username:password@hostname:port http://target.com
// curl -p --proxy administartor:1234567@127.0.0.1:7100 http://www.baidu.com
func main() {
	//HTTP处理器
	handler := http.HandlerFunc(tunnel)
	//建立HTTP服务器
	err := http.ListenAndServe(addr, handler)
	if err != nil {
		panic(err)
	}
}
